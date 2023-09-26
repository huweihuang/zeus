package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ware "github.com/huweihuang/golib/gin/middlewares"

	"github.com/huweihuang/gin-api-frame/pkg/errors"
	"github.com/huweihuang/gin-api-frame/pkg/service"
	"github.com/huweihuang/gin-api-frame/pkg/types"
	"github.com/huweihuang/gin-api-frame/pkg/validation"
)

type InstanceHandler struct {
	service *service.InstanceService
}

func newInstanceHandler() *InstanceHandler {
	return &InstanceHandler{
		service: service.NewInstanceService(),
	}
}

// 创建实例
func (h *InstanceHandler) CreateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateCreateInstance(&instance)
	if len(errs) != 0 {
		ware.ValidateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.CreateInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			ware.NotFoundWrapper(c, "instance", map[string]interface{}{"error": err.Error()})
			return
		}
		ware.ErrorWrapper(c, "CreateInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	ware.SucceedWrapper(c, "CreateInstance", data)
}

// 更新实例
func (h *InstanceHandler) UpdateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateUpdateInstance(&instance)
	if len(errs) != 0 {
		ware.ValidateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.UpdateInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			ware.NotFoundWrapper(c, "instance", map[string]interface{}{"error": err.Error()})
			return
		}
		ware.ErrorWrapper(c, "UpdateInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	ware.SucceedWrapper(c, "UpdateInstance", data)
}

// 查询实例任务创建结果
func (h *InstanceHandler) GetInstance(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		err := fmt.Errorf("name is required")
		ware.BadRequestWrapper(c, err)
		return
	}
	instance := types.Instance{}
	instance.Name = name

	e, err := h.service.GetInstance(&instance)
	if err != nil {
		if err == errors.ErrJobNotFound {
			data := map[string]string{
				"name": instance.Name,
			}
			ware.NotFoundWrapper(c, "jobID", data)
			return
		}
		ware.ErrorWrapper(c, "GetInstance", err)
		return
	}
	ware.SucceedWrapper(c, "GetInstance", e)
}

// 删除实例
func (h *InstanceHandler) DeleteInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateDeleteInstance(&instance)
	if len(errs) != 0 {
		ware.ValidateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.DeleteInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			ware.NotFoundWrapper(c, "instance", instance)
			return
		}
		ware.ErrorWrapper(c, "DeleteInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	ware.SucceedWrapper(c, "DeleteInstance", data)
}
