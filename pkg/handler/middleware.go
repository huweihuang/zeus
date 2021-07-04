package handler

import (
	"fmt"
	"net/http"

	"github.com/huweihuang/gin-api-frame/pkg/types"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/huweihuang/gin-api-frame/pkg/apis"
	"github.com/huweihuang/gin-api-frame/pkg/util/log"
)

const (
	ControllerCtx  = "controllerCtx"
	instanceReqCtx = "instanceReq"
)

// Middleware: 处理公共解析操作
func HandlerMiddleware(c *gin.Context) {
	logger := log.InitHttpLogger(c)
	logger.Debug("Use HandlerMiddleware")

	instance := types.Instance{}
	if err := c.BindJSON(&instance); err != nil {
		resp := apis.Response{
			Code:    http.StatusBadRequest,
			Message: "invalid request body",
			Data:    map[string]interface{}{"error": err},
		}
		logger.WithField("err", err).Warn("Invalid request body")
		c.JSON(http.StatusBadRequest, resp)
		return
	}
	c.Set(instanceReqCtx, instance)
}

// 封装请求成功的处理逻辑
func succeedWrapper(c *gin.Context, msg string, data interface{}) {
	logger := log.InitHttpLogger(c)
	resp := apis.Response{
		Code:    http.StatusOK,
		Message: fmt.Sprintf("%s succeed", msg),
		Data:    data,
	}
	logger.WithFields(logrus.Fields{
		"data": data,
	}).Infof("%s succeed", msg)
	c.JSON(http.StatusOK, resp)
}

// 封装请求失败的处理逻辑
func errorWrapper(c *gin.Context, msg string, err error) {
	logger := log.InitHttpLogger(c)
	resp := apis.Response{
		Code:    http.StatusInternalServerError,
		Message: fmt.Sprintf("%s failed", msg),
		Data:    map[string]interface{}{"error": err.Error()},
	}
	logger.WithFields(logrus.Fields{
		"err": err,
	}).Errorf("%s failed", msg)
	c.JSON(http.StatusInternalServerError, resp)
}

// 封装NotFound的处理逻辑
func notFoundWrapper(c *gin.Context, msg string, data interface{}) {
	logger := log.InitHttpLogger(c)
	resp := apis.Response{
		Code:    http.StatusNotFound,
		Message: fmt.Sprintf("%s not found", msg),
		Data:    data,
	}
	logger.WithFields(logrus.Fields{
		"data": data,
	}).Warnf("%s not found", msg)
	c.JSON(http.StatusNotFound, resp)
}

// 封装校验非法请求的处理逻辑
func validateBadRequestWrapper(c *gin.Context, errs field.ErrorList) {
	logger := log.InitHttpLogger(c)
	resp := apis.Response{
		Code:    http.StatusBadRequest,
		Message: "invalid request",
		Data:    map[string]interface{}{"error": errs},
	}
	logger.WithFields(logrus.Fields{
		"errs": errs.ToAggregate(),
	}).Error("Invalid request")
	c.JSON(http.StatusBadRequest, resp)
}

// 封装非法请求的处理逻辑
func badRequestWrapper(c *gin.Context, err error) {
	logger := log.InitHttpLogger(c)
	resp := apis.Response{
		Code:    http.StatusBadRequest,
		Message: "invalid request body",
		Data:    map[string]interface{}{"error": err.Error()},
	}
	logger.WithField("err", err).Error("Invalid request body")
	c.JSON(http.StatusBadRequest, resp)
}
