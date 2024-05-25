package stream

import (
	"context"
	streamChat "github.com/GetStream/stream-chat-go/v6"
	"log"
	"time"
)

var client *streamChat.Client

type Configuration struct {
	ApiKey    string `mapstructure:"api_key"`
	ApiSecret string `mapstructure:"api_secret"`
}

func New(conf Configuration) {
	_client, err := streamChat.NewClient(conf.ApiKey, conf.ApiSecret)
	if err != nil {
		log.Panicln("初始化stream chat失败", err)
	}
	client = _client
}

func Client() *streamChat.Client {
	return client
}

// CreateToken 创建令牌
func CreateToken(userID string, expire time.Time, issuedAt ...time.Time) (string, error) {
	return client.CreateToken(userID, expire, issuedAt...)
}

// UpsertUsers 插入或更新用户
func UpsertUsers(users ...*streamChat.User) (*streamChat.UsersResponse, error) {
	return client.UpsertUsers(context.Background(), users...)
}
