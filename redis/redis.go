package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"net"
	"runtime"
	"time"
)

var (
	rdb redis.UniversalClient
)

type Configuration struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	PoolSize int    `mapstructure:"pool_size"`
	DB       int    `mapstructure:"db"`
}

func New(conf Configuration) redis.UniversalClient {
	rdb = redis.NewUniversalClient(&redis.UniversalOptions{
		Addrs:    []string{net.JoinHostPort(conf.Host, conf.Port)},
		Username: conf.Username,
		Password: conf.Password,
		DB:       conf.DB,
		PoolSize: conf.PoolSize * runtime.GOMAXPROCS(0),
	})

	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return rdb
}

// Set 设置键值对
func Set(key string, value []byte, expiration time.Duration) error {
	return rdb.Set(context.Background(), key, value, expiration*time.Second).Err()
}

// Get 获取键值
func Get(key string) ([]byte, error) {
	return rdb.Get(context.Background(), key).Bytes()
}

// Del 删除键
func Del(key string) error {
	return rdb.Del(context.Background(), key).Err()
}
