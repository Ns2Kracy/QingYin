package upload

import (
	"QingYin/global"
	"errors"
	"io"
	"mime/multipart"
	"strconv"
	"time"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"go.uber.org/zap"
)

type AliyunOSS struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@object: *AliyunOSS
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*AliyunOSS) UploadFile(file *multipart.FileHeader) (string, string, error) {
	//创建Bucket实例
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	//读取本地文件
	f, openErr := file.Open()
	if openErr != nil {
		global.GVA_LOG.Error("function file.Opne() Failed", zap.Any("err", openErr.Error()))
		return "", "", errors.New("function file.Open() Failed" + openErr.Error())
	}
	defer f.Close()

	//上传阿里云OSS路径
	yunUploadPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "uploads" + "/" + time.Now().Format("2006-01-02") + "/" + file.Filename

	//上传文件流
	upErr := bucket.PutObject(yunUploadPath, f)
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
//@object: *AliyunOSS
//@function: UploadImage
//@description: 上传图片
//@param: file *multipart.FileHeader
//@return: string, string, error
func (*AliyunOSS) UploadImage(buf io.Reader) (string, string, error) {
	//创建Bucket实例
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}

	//上传阿里云OSS路径
	yunUploadPath := global.GVA_CONFIG.AliyunOSS.BasePath + "/" + "snapshots" + "/" + strconv.FormatInt(time.Now().Unix(), 10) + "_snapshot" + ".jpg"

	//上传文件流
	upErr := bucket.PutObject(yunUploadPath, buf)
	if upErr != nil {
		global.GVA_LOG.Error("function formUploader.Put() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function formUploader.Put() Failed, err:" + err.Error())
	}

	//设置请求头ContentType类型避免访问URL时直接下载
	options := oss.ContentType("image/jpg")
	err = bucket.SetObjectMeta(yunUploadPath, options)
	if err != nil {
		global.GVA_LOG.Error("function bucket.SetObjectMeta() Failed", zap.Any("err", err.Error()))
		return "", "", errors.New("function bucket.SetObjectMeta() Failed, err:" + err.Error())
	}

	return global.GVA_CONFIG.AliyunOSS.BucketUrl + "/" + yunUploadPath, yunUploadPath, nil
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@object: *AliyunOSS
//@function: DeleteFile
//@description: 删除文件
//@param: file *multipart.FileHeader
//@return: error
func (*AliyunOSS) DeleteFile(key string) error {
	//创建Bucket实例
	bucket, err := NewBucket()
	if err != nil {
		global.GVA_LOG.Error("function AliyunOSS.NewBucket() Failed", zap.Any("err", err.Error()))
		return errors.New("function AliyunOSS.NewBucket() Failed, err:" + err.Error())
	}
	//删除单个文件。objectName表示删除OSS文件时需要指定包含文件后缀在内的完整路径，例如abc/efg/123.jpg
	deErr := bucket.DeleteObject(key)
	if deErr != nil {
		global.GVA_LOG.Error("function AliyunOSS.DeleteObject() Failed", zap.Any("err", deErr.Error()))
		return errors.New("function AliyunOSS.DeleteObject() Failed, err:" + deErr.Error())
	}
	return nil
}

//@author [Mr-Jacks520](https://github.com/Mr-Jacks520)
//@function: NewBucket
//@description: 创建Bucket实例
//@return: *oss.Bucket,error
func NewBucket() (*oss.Bucket, error) {
	// 创建OSSClient实例。
	client, err := oss.New(global.GVA_CONFIG.AliyunOSS.Endpoint, global.GVA_CONFIG.AliyunOSS.AccessKeyId, global.GVA_CONFIG.AliyunOSS.AccessKeySecret)
	if err != nil {
		return nil, err
	}

	// 获取存储空间。
	bucket, err := client.Bucket(global.GVA_CONFIG.AliyunOSS.BucketName)
	if err != nil {
		return nil, err
	}

	return bucket, nil
}
