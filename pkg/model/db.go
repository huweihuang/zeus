package model

import (
	"github.com/huweihuang/golib/db"
	"gorm.io/gorm"

	"github.com/huweihuang/zeus/cmd/server/app/configs"
)

// DB is a global db manager
var DB *DBMng

type DBMng struct {
	db *gorm.DB
}

func NewDBMng(db *gorm.DB) *DBMng {
	return &DBMng{
		db: db,
	}
}

func InitDB(dbConf *configs.DBConfig) (*DBMng, error) {
	d, err := db.SetupDB(dbConf.Addr, dbConf.DBName, dbConf.User, dbConf.Password, dbConf.LogLevel)
	if err != nil {
		return nil, err
	}
	DB = NewDBMng(d)
	return DB, nil
}

func Close() error {
	db, err := DB.db.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
