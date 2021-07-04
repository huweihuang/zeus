package model

import (
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var engine *gorm.DB

// 数据库初始化
func SetupDB(dsn string) (err error) {
	engine, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	db, err := engine.DB()
	if err != nil {
		return err
	}
	// Set the maximum lifetime of the connection (less than server setting)
	db.SetConnMaxLifetime(30 * time.Minute)

	return nil
}

// 获取数据库引擎
func GetDB() *gorm.DB {
	return engine
}

// 关闭数据库
func Close() error {
	db, err := engine.DB()
	if err != nil {
		return err
	}
	return db.Close()
}
