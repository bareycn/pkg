package db

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

var db *gorm.DB

type Configuration struct {
	Type         string          `mapstructure:"type"`
	Dsn          string          `mapstructure:"dsn"`
	MaxIdleConns int             `mapstructure:"max_idle_conns"`
	MaxOpenConns int             `mapstructure:"max_open_conns"`
	Logger       logger.LogLevel `mapstructure:"logger"`
}

func New(conf Configuration) {
	switch conf.Type {
	case "mysql":
		db = NewMysql(conf)
	case "postgres":
		db = NewPostgres(conf)
	default:
		db = NewMysql(conf)
	}
}

func DB() *gorm.DB {
	return db
}

func Model(value interface{}) *gorm.DB {
	if db == nil {
		log.Panicln("数据库未初始化")
	}
	return db.Model(value)
}

// Transaction 事务
func Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error) {
	return db.Transaction(fc, opts...)
}
