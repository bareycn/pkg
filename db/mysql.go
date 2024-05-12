package db

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewMysql(conf Configuration) *gorm.DB {
	if _db, err := gorm.Open(mysql.New(mysql.Config{DSN: conf.Dsn}), &gorm.Config{
		Logger: logger.Default.LogMode(conf.Logger),
	}); err != nil {
		panic(err)
	} else {
		sqlDB, _ := _db.DB()
		sqlDB.SetMaxIdleConns(conf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(conf.MaxOpenConns)
		return _db
	}
}
