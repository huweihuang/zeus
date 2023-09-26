package model

import (
	"github.com/huweihuang/golib/db"
	"gorm.io/gorm"

	"github.com/huweihuang/zeus/cmd/server/app/configs"
)

var DB *gorm.DB

func SetupDB(dbConf *configs.DBConfig) (DB *gorm.DB, err error) {
	DB, err = db.SetupDB(dbConf.Addr, dbConf.DBName, dbConf.User, dbConf.Password, dbConf.LogLevel)
	if err != nil {
		return nil, err
	}
	return DB, nil
}

func Close() error {
	sqlDB, err := DB.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
