package controller

import (
	"github.com/huweihuang/gin-api-frame/pkg/model"
	"github.com/huweihuang/gin-api-frame/pkg/types"
)

// 控制器用来批量管理服务
type InstanceController struct {
}

// 初始化一个控制器
func NewInstanceController() InstanceInterface {
	return &InstanceController{}
}

// 创建实例
func (c *InstanceController) CreateInstance(ins *types.Instance) error {
	err := model.CreateInstance(ins)
	if err != nil {
		return err
	}
	return nil
}

// 查询实例
func (c *InstanceController) GetInstance(ins *types.Instance) (*types.Instance, error) {
	ins, err := model.GetInstanceByHostIDAndAppName(ins.Spec.HostID, ins.Name)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新实例
func (c *InstanceController) UpdateInstance(ins *types.Instance) error {
	err := model.UpdateInstanceImage(ins.Name, ins.Spec.Image)
	if err != nil {
		return err
	}
	return nil
}

// 删除实例
func (c *InstanceController) DeleteInstance(ins *types.Instance) error {
	err := model.DeleteInstance(ins.Name)
	if err != nil {
		return err
	}
	return nil
}
