package server

import (
	"github.com/huweihuang/gin-api-frame/pkg/handlers"
	log "github.com/huweihuang/golib/logger/logrus"
)

func (s *Server) setupRoutes() {
	group := s.gin.Group("/api/v1")
	handler := handlers.New()
	// instance
	group.POST("/instance", handlers.HandlerMiddleware, handler.CreateInstance)
	group.PUT("/instance", handlers.HandlerMiddleware, handler.UpdateInstance)
	group.GET("/instance", handler.GetInstance)
	group.DELETE("/instance", handler.DeleteInstance)

	log.Logger.Info("setup routes succeed")
}
