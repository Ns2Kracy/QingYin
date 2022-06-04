package service

import (
	"QingYin/global"
	model "QingYin/model/system"
	"QingYin/utils/upload"
	"bytes"
	"fmt"
	"mime/multipart"
	"os"

	ffmpeg "github.com/u2takey/ffmpeg-go"
	"go.uber.org/zap"
)

//针对接口:
// /douyin/publish/list
// /douyin/publish/action

type PublishService struct{}

//投稿业务逻辑
func (p *PublishService) Action(file *multipart.FileHeader, video *model.SysVideo) error {
	//文件上传
	oss := upload.NewOss()
	playURL, _, upErr := oss.UploadFile(file)
	if upErr != nil {
		global.GVA_LOG.Error("Action Failed", zap.Any("err", upErr.Error()))
		return upErr
	}
	video.PlayURL = "https://" + playURL                  //设置播放地址
	coverURL, Generr := p.generateCoverURL(video.PlayURL) //生成缩略图
	if Generr != nil {
		global.GVA_LOG.Error("Generate CoverURL Failed", zap.Any("err", Generr.Error()))
		return Generr
	}
	video.CoverURL = coverURL //设置视频封面地址
	err := global.GVA_DB.Create(&video).Error
	return err
}

//生成并上传视频缩略图
func (*PublishService) generateCoverURL(videoPath string) (string, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", 1)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		global.GVA_LOG.Error("Transfer Image Failed", zap.Any("err", err.Error()))
		return "", err
	}

	// img, err := imaging.Decode(buf)
	// if err != nil {
	// 	return "", err
	// }

	//图片上传
	oss := upload.NewOss()
	coverURL, _, upErr := oss.UploadImage(buf)
	if upErr != nil {
		global.GVA_LOG.Error("Upload Image Failed", zap.Any("err", upErr.Error()))
		return "", upErr
	}
	return "https://" + coverURL, nil

}
