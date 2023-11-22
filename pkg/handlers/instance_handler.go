package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	ware "github.com/huweihuang/golib/gin/middlewares"

	"github.com/huweihuang/zeus/pkg/errors"
	"github.com/huweihuang/zeus/pkg/service"
	"github.com/huweihuang/zeus/pkg/types"
	"github.com/huweihuang/zeus/pkg/validation"
)

type InstanceHandler struct {
	service *service.InstanceService
}

func newInstanceHandler() *InstanceHandler {
	return &InstanceHandler{
		service: service.NewInstanceService(),
	}
}

// CreateInstance godoc
//
//	@Summary		Create Instance
//	@Description	Create Instance.
//	@Tags			Instance
//	@Accept			json
//	@Produce		json
//	@Param			request	body	types.Instance	true	"Request body"
//	@Produce		json
//	@Success		200	{object}	types.CreateInstanceResp
//	@Failure		400	{object}	types.ErrorResp
//	@Failure		500	{object}	types.ErrorResp
//	@Router			/api/v1/instance [POST]
func (h *InstanceHandler) CreateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	err := validation.ValidateCreateInstance(&instance)
	if err != nil {
		ware.BadRequestWrapper(c, err)
		return
	}

	err = h.service.CreateInstance(&instance)
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

// UpdateInstance godoc
//
//	@Summary		Update Instance
//	@Description	Update Instance.
//	@Tags			Instance
//	@Accept			json
//	@Produce		json
//	@Param			request	body	types.Instance	true	"Request body"
//	@Produce		json
//	@Success		200	{object}	types.CreateInstanceResp
//	@Failure		400	{object}	types.ErrorResp
//	@Failure		500	{object}	types.ErrorResp
//	@Router			/api/v1/instance [PUT]
func (h *InstanceHandler) UpdateInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	err := validation.ValidateUpdateInstance(&instance)
	if err != nil {
		ware.BadRequestWrapper(c, err)
		return
	}

	err = h.service.UpdateInstance(&instance)
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

// GetInstance godoc
//
//	@Summary		Get Instance
//	@Description	Get Instance.
//	@Tags			Instance
//	@Accept			json
//	@Produce		json
//	@Param			name	query	string	true	"name"
//	@Produce		json
//	@Success		200	{object}	types.GetInstanceResp
//	@Failure		400	{object}	types.ErrorResp
//	@Failure		500	{object}	types.ErrorResp
//	@Router			/api/v1/instance [GET]
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

// DeleteInstance godoc
//
//	@Summary		Delete Instance
//	@Description	Delete Instance.
//	@Tags			Instance
//	@Accept			json
//	@Produce		json
//	@Param			request	body	types.Instance	true	"Request body"
//	@Produce		json
//	@Success		200	{object}	types.CreateInstanceResp
//	@Failure		400	{object}	types.ErrorResp
//	@Failure		500	{object}	types.ErrorResp
//	@Router			/api/v1/instance [DELETE]
func (h *InstanceHandler) DeleteInstance(c *gin.Context) {
	instance := c.MustGet(instanceReqCtx).(types.Instance)
	err := validation.ValidateDeleteInstance(&instance)
	if err != nil {
		ware.BadRequestWrapper(c, err)
		return
	}

	err = h.service.DeleteInstance(&instance)
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
