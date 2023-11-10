package service

import (
	"github.com/huweihuang/zeus/pkg/model"
	"github.com/huweihuang/zeus/pkg/types"
)

// 控制器用来批量管理服务
type InstanceService struct {
	db *model.DBMng
}

// 初始化一个控制器
func NewInstanceService() *InstanceService {
	return &InstanceService{
		db: model.DB,
	}
}

// 创建实例
func (s *InstanceService) CreateInstance(ins *types.Instance) error {
	err := s.db.CreateInstance(ins)
	if err != nil {
		return err
	}
	return nil
}

// 查询实例
func (s *InstanceService) GetInstance(ins *types.Instance) (*types.Instance, error) {
	ins, err := s.db.GetInstance(ins.Spec.HostID, ins.Name)
	if err != nil {
		return nil, err
	}
	return ins, nil
}

// 更新实例
func (s *InstanceService) UpdateInstance(ins *types.Instance) error {
	err := s.db.UpdateInstanceImage(ins.Name, ins.Spec.Image)
	if err != nil {
		return err
	}
	return nil
}

// 删除实例
func (s *InstanceService) DeleteInstance(ins *types.Instance) error {
	err := s.db.DeleteInstance(ins.Name)
	if err != nil {
		return err
	}
	return nil
}
