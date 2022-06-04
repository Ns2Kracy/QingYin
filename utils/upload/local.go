package upload

import (
	"QingYin/global"
	"QingYin/utils"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"path"
	"strings"
	"time"

	"go.uber.org/zap"
)

//本地上传
type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*Local) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//读取文件后缀
	ext := path.Ext(file.Filename)
	//读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	//拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext
	//尝试创建上传路径
	if err := os.Mkdir(global.GVA_CONFIG.Local.Path, os.ModePerm); err != nil {
		global.GVA_LOG.Error("function os.MkdirAll() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function os.MkdirAll() Filed, err:" + err.Error())
	}

	//拼接上传路径以及文件名
	p := global.GVA_CONFIG.Local.Path + "/" + filename

	//读取文件
	f, openErr := file.Open()
	if openErr != nil {
		global.GVA_LOG.Error("function file.Open() Failed", zap.Any("err", openErr.Error()))
		return "", "", errors.New("function file.Open() Filed, err:" + openErr.Error())
	}
	defer f.Close() //关闭文件

	//创建文件
	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() Failed", zap.Any("err", createErr.Error()))
		return "", "", errors.New("function os.Create() Filed" + createErr.Error())
	}
	defer out.Close()

	//传输拷贝文件
	_, copyErr := io.Copy(out, f)
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() Failed", zap.Any("err", copyErr.Error()))
		return "", "", errors.New("function io.Copy() Filed" + copyErr.Error())
	}
	return p, filename, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@object: *Local
//@function: UploadImage
//@description: 上传图片>>>>>>>>>>>>>>>>>>未实现
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*Local) UploadImage(buf io.Reader) (string, string, error) {
	//创建Bucket实例
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	//上传阿里云OSS路径
	yunUploadPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "snapshots" + "/" + time.Now().Format("2006-01-02") + "_snapshot" + ".jpg"

	//上传文件流
	upErr := bucket.PutObject(yunUploadPath, buf)
	if upErr != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunUploadPath, yunUploadPath, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: file *multipart.FileHeader
//@return: error
func (*Local) DeleteFile(key string) error {
	p := global.GVA_CONFIG.Local.Path + "/" + key
	if strings.Contains(p, global.GVA_CONFIG.Local.Path) {
		if err := os.Remove(p); err != nil {
			return errors.New("本地文件删除失败, err:" + err.Error())
		}
	}
	return nil
}
