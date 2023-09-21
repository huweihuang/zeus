package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/huweihuang/gin-api-frame/pkg/errors"
	"github.com/huweihuang/gin-api-frame/pkg/service"
	"github.com/huweihuang/gin-api-frame/pkg/types"
	"github.com/huweihuang/gin-api-frame/pkg/validation"
)

// 创建实例
func CreateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateCreateInstance(&instance)
	if len(errs) != 0 {
		validateBadRequestWrapper(c, errs)
		return
	}

	ic := c.MustGet(ControllerCtx).(service.InstanceInterface)
	err := ic.CreateInstance(&instance)
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
func UpdateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateUpdateInstance(&instance)
	if len(errs) != 0 {
		validateBadRequestWrapper(c, errs)
		return
	}

	ic := c.MustGet(ControllerCtx).(service.InstanceInterface)
	err := ic.UpdateInstance(&instance)
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
func GetInstance(c *gin.Context) {
	name := c.Query("name")

	if name == "" {
		err := fmt.Errorf("name is required")
		badRequestWrapper(c, err)
		return
	}
	instance := types.Instance{}
	instance.Name = name

	ic := c.MustGet(ControllerCtx).(service.InstanceInterface)
	e, err := ic.GetInstance(&instance)
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
func DeleteInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	errs := validation.ValidateDeleteInstance(&instance)
	if len(errs) != 0 {
		validateBadRequestWrapper(c, errs)
		return
	}

	ic := c.MustGet(ControllerCtx).(service.InstanceInterface)
	err := ic.DeleteInstance(&instance)
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
