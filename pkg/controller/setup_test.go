package controller_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/pkg/controller"
	"github.com/huweihuang/gin-api-frame/pkg/model"
	"github.com/huweihuang/gin-api-frame/pkg/server"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

var (
	InsCtrl controller.InstanceInterface
)

func init() {
	log.InitLogger("", "debug", "", "text", true)
}

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	os.Exit(ret)
}

func setup() {
	conf := config.MustLoad(os.Getenv("TEST_GIN_CONFIG"))
	if err := setupDB(&conf.Database); err != nil {
		panic(err)
	}
	InsCtrl = server.InsCtrl
}

func setupDB(dbConf *config.DBConfig) error {
	if err := model.SetupDB(config.FormatDSN(dbConf)); err != nil {
		return fmt.Errorf("failed to setup database")
	}
	return nil
}
