package server

import (
	log "github.com/huweihuang/golib/logger/zap"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/huweihuang/zeus/docs"
	"github.com/huweihuang/zeus/pkg/handlers"
)

func (s *Server) setupRoutes() {
	// http://127.0.0.1/swagger/index.html
	s.gin.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	group := s.gin.Group("/api/v1")
	handler := handlers.New()
	// instance
	group.POST("/instance", handlers.HandlerMiddleware, handler.CreateInstance)
	group.PUT("/instance", handlers.HandlerMiddleware, handler.UpdateInstance)
	group.GET("/instance", handler.GetInstance)
	group.DELETE("/instance", handler.DeleteInstance)

	log.Logger().Info("setup routes succeed")
}
