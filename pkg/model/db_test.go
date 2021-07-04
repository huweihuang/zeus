package model

import (
	"os"
	"testing"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

func init() {
	log.InitLogger("", "debug", "", "text", true)
}

func TestSetupDB(t *testing.T) {
	conf := config.MustLoad(os.Getenv("TEST_GIN_CONFIG"))
	dsn := config.FormatDSN(&conf.Database)
	if err := SetupDB(dsn); err != nil {
		t.Errorf("Failed to setup db: %s", err)
	}
}
