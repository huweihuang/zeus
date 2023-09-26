package service

import (
	"github.com/huweihuang/zeus/pkg/model"
	"github.com/huweihuang/zeus/pkg/types"
)

// 控制器用来批量管理服务
type InstanceService struct {
}

// 初始化一个控制器
func NewInstanceService() *InstanceService {
	return &InstanceService{}
}

// 创建实例
func (c *InstanceService) CreateInstance(ins *types.Instance) error {
	err := model.CreateInstance(ins)
	if err != nil {
		return err
	}
	return nil
}

// 查询实例
func (c *InstanceService) GetInstance(ins *types.Instance) (*types.Instance, error) {
	ins, err := model.GetInstanceByHostIDAndAppName(ins.Spec.HostID, ins.Name)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新实例
func (c *InstanceService) UpdateInstance(ins *types.Instance) error {
	err := model.UpdateInstanceImage(ins.Name, ins.Spec.Image)
	if err != nil {
		return err
	}
	return nil
}

// 删除实例
func (c *InstanceService) DeleteInstance(ins *types.Instance) error {
	err := model.DeleteInstance(ins.Name)
	if err != nil {
		return err
	}
	return nil
}
