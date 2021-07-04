package handler_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"

	"github.com/huweihuang/gin-api-frame/cmd/server/app/config"
	"github.com/huweihuang/gin-api-frame/pkg/model"
	"github.com/huweihuang/gin-api-frame/pkg/server"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
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
	conf := config.MustLoad(os.Getenv("TEST_CONFIG"))
	if err := setupDB(&conf.Database); err != nil {
		panic(err)
	}
}

func setupDB(dbConf *config.DBConfig) error {
	if err := model.SetupDB(config.FormatDSN(dbConf)); err != nil {
		return fmt.Errorf("failed to setup database")
	}
	return nil
}

func ginTest(req *http.Request) (*gin.Context, *gin.Engine, *httptest.ResponseRecorder) {
	w := httptest.NewRecorder()
	gin.SetMode(gin.ReleaseMode)

	ctx, engine := gin.CreateTestContext(w)
	engine.Use(server.RegisterController)

	ctx.Request = req
	return ctx, engine, w
}
