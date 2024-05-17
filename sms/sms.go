package sms

import "errors"

type Configuration struct {
	Type   string               `mapstructure:"type"`
	Aliyun *AliyunConfiguration `mapstructure:"aliyun"`
}

type Sms interface {
	Send(phone, signName, template string, data map[string]interface{}) error
}

func New(config Configuration) (Sms, error) {
	switch config.Type {
	case "aliyun":
		return NewAliyun(config.Aliyun)
	default:
		return nil, errors.New("未知的短信服务提供者")
	}
}
