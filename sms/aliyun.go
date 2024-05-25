package sms

import (
	"encoding/json"
	"errors"
	openapi "github.com/alibabacloud-go/darabonba-openapi/v2/client"
	dysmsapi "github.com/alibabacloud-go/dysmsapi-20180501/v2/client"
	"github.com/alibabacloud-go/tea/tea"
	"github.com/bareycn/pkg/rand"
	"github.com/bareycn/pkg/tpl"
	"log"
)

type AliyunConfiguration struct {
	AccessKeyId     string `mapstructure:"access_id"`
	AccessKeySecret string `mapstructure:"access_secret"`
	Endpoint        string `mapstructure:"endpoint"`  // 短信服务节点
	IsGlobe         bool   `mapstructure:"is_globe"`  // 是否是国际短信
	SenderID        string `mapstructure:"sender_id"` // 国际短信需要发送者ID
}

type AliyunProvider struct {
	client *dysmsapi.Client
}

func NewAliyun(conf *AliyunConfiguration) Sms {
	var err error
	client, err := dysmsapi.NewClient(&openapi.Config{
		AccessKeyId:     tea.String(conf.AccessKeyId),
		AccessKeySecret: tea.String(conf.AccessKeySecret),
		Endpoint:        tea.String(conf.Endpoint),
	})
	if err != nil {
		log.Panicln(err)
	}
	return &AliyunProvider{
		client: client,
	}
}

// Send 发送国际短信,完整消息内容
func (a *AliyunProvider) Send(phone, signName, template string, data map[string]interface{}) error {
	taskID := rand.String(32)
	parseBody, err := tpl.ParseTemplate(template, template, data)
	if err != nil {
		return err
	}
	result, err := a.client.SendMessageToGlobe(&dysmsapi.SendMessageToGlobeRequest{
		To:      tea.String(phone),
		Message: tea.String(parseBody),
		From:    tea.String(signName),
		TaskId:  tea.String(taskID),
	})
	if err != nil {
		return err
	}
	if result.Body.ResponseCode != tea.String("OK") {
		return errors.New(tea.StringValue(result.Body.ResponseCode))
	}
	return nil
}

// SendTemplate 发送模板短信
func (a *AliyunProvider) SendTemplate(phone, signName, template string, data map[string]interface{}) error {
	paramsStr, err := json.Marshal(data)
	if err != nil {
		return err
	}
	result, err := a.client.SendMessageWithTemplate(&dysmsapi.SendMessageWithTemplateRequest{
		To:            tea.String(phone),
		From:          tea.String(signName),
		TemplateCode:  tea.String(template),
		TemplateParam: tea.String(string(paramsStr)),
	})
	if err != nil {
		return err
	}
	if result.Body.ResponseCode != tea.String("OK") {
		return errors.New(tea.StringValue(result.Body.ResponseCode))
	}
	return nil
}
