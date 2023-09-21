package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"

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
		validateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.CreateInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			notFoundWrapper(c, "instance", map[string]interface{}{"error": err.Error()})
			return
		}
		errorWrapper(c, "CreateInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	succeedWrapper(c, "CreateInstance", data)
}

// 更新实例
func (h *InstanceHandler) UpdateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateUpdateInstance(&instance)
	if len(errs) != 0 {
		validateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.UpdateInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			notFoundWrapper(c, "instance", map[string]interface{}{"error": err.Error()})
			return
		}
		errorWrapper(c, "UpdateInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	succeedWrapper(c, "UpdateInstance", data)
}

// 查询实例任务创建结果
func (h *InstanceHandler) GetInstance(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		err := fmt.Errorf("name is required")
		badRequestWrapper(c, err)
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
			notFoundWrapper(c, "jobID", data)
			return
		}
		errorWrapper(c, "GetInstance", err)
		return
	}
	succeedWrapper(c, "GetInstance", e)
}

// 删除实例
func (h *InstanceHandler) DeleteInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateDeleteInstance(&instance)
	if len(errs) != 0 {
		validateBadRequestWrapper(c, errs)
		return
	}

	err := h.service.DeleteInstance(&instance)
	if err != nil {
		if err == errors.ErrInstanceNotFound {
			notFoundWrapper(c, "instance", instance)
			return
		}
		errorWrapper(c, "DeleteInstance", err)
		return
	}
	data := map[string]string{"jobID": instance.JobID}
	succeedWrapper(c, "DeleteInstance", data)
}
