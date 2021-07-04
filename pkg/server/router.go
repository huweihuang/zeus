package server

import (
	"github.com/gin-gonic/gin"

	"github.com/huweihuang/gin-api-frame/pkg/controller"
	"github.com/huweihuang/gin-api-frame/pkg/handler"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

var (
	InsCtrl controller.InstanceInterface
)

func (s *apiserver) setupRoutes() {
	s.gin.Use(RegisterController)
	group := s.gin.Group("/api/v1")
	// instance
	group.POST("/instance", handler.HandlerMiddleware, handler.CreateInstance)
	group.PUT("/instance", handler.HandlerMiddleware, handler.UpdateInstance)
	group.GET("/instance", handler.GetInstance)
	group.DELETE("/instance", handler.DeleteInstance)

	log.Logger.Info("setup routes succeed")
}

// 注册controller
func RegisterController(c *gin.Context) {
	logger := log.InitHttpLogger(c)
	logger.Debug("register controller")
	InsCtrl = controller.NewInstanceController()
	c.Set(handler.ControllerCtx, InsCtrl)
}
