package storage

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
	"time"
)

type OssProvider struct {
	client *oss.Client
	bucket *oss.Bucket
}

type OssConfiguration struct {
	AccessID     string `mapstructure:"access_id"`
	AccessSecret string `mapstructure:"access_secret"`
	Endpoint     string `mapstructure:"endpoint"`
	Bucket       string `mapstructure:"bucket"`
}

// NewOss 初始化OSS
func NewOss(conf OssConfiguration) Storage {
	c, err := oss.New(conf.Endpoint, conf.AccessID, conf.AccessSecret)
	if err != nil {
		log.Panicf("OSS Error: %s", err.Error())
	}
	b, err := c.Bucket(conf.Bucket)
	if err != nil {
		log.Panicf("Bucket Error: %s", err.Error())
	}
	return &OssProvider{
		client: c,
		bucket: b,
	}
}

// PutFileKey 上传文件
func (o *OssProvider) PutFileKey(key string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	err = o.bucket.PutObject(key, src)
	return key, err
}

// PutFile 上传文件
func (o *OssProvider) PutFile(path string, file *multipart.FileHeader) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()
	ext := filepath.Ext(file.Filename)
	fileKey := fmt.Sprintf("%s/%s%s", path, time.Now().Format("2006/01/02")+uuid.NewString(), ext)
	err = o.bucket.PutObject(fileKey, src)
	return fileKey, err
}

// PutObject 上传文件
func (o *OssProvider) PutObject(key string, reader io.Reader) (string, error) {
	return key, o.bucket.PutObject(key, reader)
}
