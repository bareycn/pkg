package storage

import (
	"io"
	"log"
	"mime/multipart"
)

type Storage interface {
	PutFile(path string, file *multipart.FileHeader) (string, error)
	PutObject(key string, reader io.Reader) (string, error)
}

type Configuration struct {
	Provider string           `mapstructure:"provider"`
	Domain   string           `mapstructure:"domain"`
	Oss      OssConfiguration `mapstructure:"oss"`
}

var storage Storage

func New(conf Configuration) {
	switch conf.Provider {
	case "oss":
		storage = NewOss(conf.Oss)
	default:
		log.Panicln("未知存储类型")
	}
}

func Client() Storage {
	return storage
}
