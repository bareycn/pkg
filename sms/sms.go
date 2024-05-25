package sms

import (
	"errors"
	"log"
)

type Configuration struct {
	Type   string               `mapstructure:"type"`
	Aliyun *AliyunConfiguration `mapstructure:"aliyun"`
}

type Sms interface {
	Send(phone, signName, template string, data map[string]interface{}) error
	SendTemplate(phone, signName, template string, data map[string]interface{}) error
}

var smsProvider Sms

func New(config Configuration) {
	switch config.Type {
	case "aliyun":
		smsProvider = NewAliyun(config.Aliyun)
	default:
		log.Panicln(errors.New("未知短信服务商"))
	}
}

func Client() Sms {
	return smsProvider
}
