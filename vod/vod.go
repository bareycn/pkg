package vod

import (
	"encoding/base64"
	"encoding/json"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	vod "github.com/alibabacloud-go/vod-20170321/v3/client"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/google/uuid"
	"io"
	"log"
	"mime/multipart"
	"path/filepath"
)

type Configuration struct {
	AccessId        string  `mapstructure:"access_id"`
	AccessSecret    string  `mapstructure:"access_secret"`
	Endpoint        string  `mapstructure:"endpoint"`
	TemplateGroupId *string `mapstructure:"template_group_id"`
	WorkflowId      *string `mapstructure:"workflow_id"`
	RegionId        string  `mapstructure:"region_id"`
	Domain          string  `mapstructure:"domain"`
}

var (
	client *vod.Client
	config *Configuration
)

type UploadAuth struct {
	AccessKeyId     string
	AccessKeySecret string
	SecurityToken   string
}

type UploadAddress struct {
	Endpoint string
	Bucket   string
	FileName string
}

// New 阿里云点播
func New(conf Configuration) {
	var err error
	client, err = vod.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(conf.AccessId),
		AccessKeySecret: tea.String(conf.AccessSecret),
		Endpoint:        tea.String(conf.Endpoint),
		RegionId:        tea.String(conf.RegionId),
	})
	config = &conf
	if err != nil {
		log.Panicln(err)
	}
}

// 解析uploadVideoResponse
func parseUploadVideoResponse(uploadVideoResponse *vod.CreateUploadVideoResponse) (*UploadAuth, *UploadAddress, error) {
	var uploadAuth *UploadAuth
	var uploadAddress *UploadAddress
	authDecode, err := base64.StdEncoding.DecodeString(tea.StringValue(uploadVideoResponse.Body.UploadAuth))
	addressDecode, err := base64.StdEncoding.DecodeString(tea.StringValue(uploadVideoResponse.Body.UploadAddress))
	if err = json.Unmarshal(authDecode, &uploadAuth); err != nil {
		return nil, nil, err
	}
	if err = json.Unmarshal(addressDecode, &uploadAddress); err != nil {
		return nil, nil, err
	}
	return uploadAuth, uploadAddress, nil
}

func initOss(auth *UploadAuth, address *UploadAddress) (*oss.Client, error) {
	return oss.New(address.Endpoint, auth.AccessKeyId, auth.AccessKeySecret, oss.SecurityToken(auth.SecurityToken))
}

func uploadOss(auth *UploadAuth, address *UploadAddress, reader io.Reader) error {
	o, err := initOss(auth, address)
	if err != nil {
		log.Panicln(err)
	}
	bucket, err := o.Bucket(address.Bucket)
	if err != nil {
		return err
	}
	err = bucket.PutObject(address.FileName, reader)
	if err != nil {
		return err
	}
	return nil
}

func UploadVideo(filename string, file multipart.File) (fileName string, mediaID string, err error) {
	title := filepath.Base(filename)
	if len(title) >= 128 {
		title = uuid.New().String()
	}

	requestUploadVideo := &vod.CreateUploadVideoRequest{
		Title:           tea.String(title),
		FileName:        tea.String(filename),
		TemplateGroupId: config.TemplateGroupId,
		WorkflowId:      config.WorkflowId,
	}
	response, err := client.CreateUploadVideo(requestUploadVideo)
	if err != nil {
		log.Panicln(err)
	}
	auth, address, err := parseUploadVideoResponse(response)
	if err != nil {
		log.Panicln(err)
	}
	err = uploadOss(auth, address, file)
	if err != nil {
		return "", "", err
	}
	return address.FileName, tea.StringValue(response.Body.VideoId), nil
}

// GetPlayInfo 获取点播地址
func GetPlayInfo(videoId string) (coverUrl string, videoUrl string, err error) {
	requestGetPlayInfo := &vod.GetPlayInfoRequest{
		VideoId: tea.String(videoId),
		Formats: tea.String("m3u8,mp4"),
	}
	response, err := client.GetPlayInfo(requestGetPlayInfo)
	if err != nil {
		return
	}
	coverUrl = tea.StringValue(response.Body.VideoBase.CoverURL)
	videoUrl = tea.StringValue(response.Body.PlayInfoList.PlayInfo[0].PlayURL)
	return
}

// GetVideoPlayAuth 点播播放凭证
func GetVideoPlayAuth(videoId string) (playAuth string, coverUrl string, err error) {
	requestGetVideoPlayAuth := &vod.GetVideoPlayAuthRequest{
		VideoId:         tea.String(videoId),
		AuthInfoTimeout: tea.Int64(60 * 60),
	}
	response, err := client.GetVideoPlayAuth(requestGetVideoPlayAuth)
	if err != nil {
		return "", "", err
	}
	return tea.StringValue(response.Body.PlayAuth), tea.StringValue(response.Body.VideoMeta.CoverURL), nil
}
