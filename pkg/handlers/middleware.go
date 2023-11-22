package handlers

import (
	"github.com/gin-gonic/gin"
	ware "github.com/huweihuang/golib/gin/middlewares"
	log "github.com/huweihuang/golib/logger/zap"

	"github.com/huweihuang/zeus/pkg/types"
)

const (
	instanceReqCtx = "instanceReq"
)

// Middleware: 处理公共解析操作
func HandlerMiddleware(c *gin.Context) {
	log.Logger().Debug("Use HandlerMiddleware")
	instance := types.Instance{}
	if err := c.BindJSON(&instance); err != nil {
		ware.BadRequestWrapper(c, err)
		return
	}
	c.Set(instanceReqCtx, instance)
}
