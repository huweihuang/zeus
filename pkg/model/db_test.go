package model

import (
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huweihuang/golib/db"
)

var (
	mockDB *DBMng
	mock   sqlmock.Sqlmock
)

func init() {
	gormDB, sqlMock, err := db.GetDBMock()
	if err != nil {
		panic(err)
	}
	mockDB = NewDBMng(gormDB)
	mock = sqlMock
}
