package mail

import (
	"crypto/tls"
	"gopkg.in/gomail.v2"
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
