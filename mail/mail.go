package mail

import (
	"bytes"
	"crypto/tls"
	"gopkg.in/gomail.v2"
	"html/template"
	"log"
)

var (
	dialer *gomail.Dialer
	config *Configuration
)

type Configuration struct {
	Name     string `mapstructure:"name"`
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}

func New(conf Configuration) {
	config = &conf
	dialer = gomail.NewDialer(conf.Host, conf.Port, conf.Username, conf.Password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}
}

// Send 发送邮件。
func Send(to, subject, body string) error {
	if dialer == nil {
		log.Panic("mail not initialized")
	}
	message := gomail.NewMessage()
	message.SetHeader("From", message.FormatAddress(config.Username, config.Name))
	message.SetHeader("To", to)
	message.SetHeader("Subject", subject)
	message.SetBody("text/html", body)

	return dialer.DialAndSend(message)
}

// ParseTemplate 解析模板字符串并返回结果字符串。
// name 为模板名称，templateStr 为模板字符串，data 为模板数据。
func ParseTemplate(name, templateStr string, data interface{}) (string, error) {
	t, err := template.New(name).Parse(templateStr)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// ParseTemplateFile 解析模板文件并返回结果字符串。
func ParseTemplateFile(templateFileName string, data interface{}) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}
	buf := new(bytes.Buffer)
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}
