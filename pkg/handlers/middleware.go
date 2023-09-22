package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	log "github.com/huweihuang/golib/logger/logrus"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/huweihuang/gin-api-frame/pkg/types"
)

const (
	instanceReqCtx = "instanceReq"
)

// Middleware: 处理公共解析操作
func HandlerMiddleware(c *gin.Context) {
	log.Logger.Debug("Use HandlerMiddleware")
	instance := types.Instance{}
	if err := c.BindJSON(&instance); err != nil {
		resp := types.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
			Data:    map[string]interface{}{"error": err},
		}
		log.Logger.WithField("err", err).Warn("Invalid request body")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.Set(instanceReqCtx, instance)
}

// 封装校验非法请求的处理逻辑
func validateBadRequestWrapper(c *gin.Context, errs field.ErrorList) {
	resp := types.Response{
		Code:    http.StatusBadRequest,
		Message: "invalid request",
		Data:    map[string]interface{}{"error": errs},
	}
	log.Logger.WithFields(logrus.Fields{
		"errs": errs.ToAggregate().Error(),
	}).Error("Invalid request")
	c.JSON(http.StatusBadRequest, resp)
}
