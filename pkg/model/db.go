package model

import (
	"github.com/huweihuang/golib/db"
	"gorm.io/gorm"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
)

var DB *gorm.DB

func SetupDB(dbConf *config.DBConfig) (*gorm.DB, error) {
	DB, err := db.SetupDB(dbConf.Addr, dbConf.DBName, dbConf.User, dbConf.Password, dbConf.LogLevel)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

// 获取数据库引擎
func GetDB() *gorm.DB {
	return DB
}

// 关闭数据库
func Close() error {
	db, err := DB.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
