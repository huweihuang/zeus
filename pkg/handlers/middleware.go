package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
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
		resp := types.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
			Data:    map[string]interface{}{"error": err},
		}
		log.Logger().With("err", err).Warn("Invalid request body")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.Set(instanceReqCtx, instance)
}
