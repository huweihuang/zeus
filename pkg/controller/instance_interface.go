package controller

import (
	"github.com/huweihuang/gin-api-frame/pkg/types"
)

// InstanceInterface is an interface for providing instance operation
type InstanceInterface interface {
	// Create instance
	CreateInstance(ins *types.Instance) error
	// Get instance status
	GetInstance(ins *types.Instance) (*types.Instance, error)
	// Update instance
	UpdateInstance(ins *types.Instance) error
	// Delete instance
	DeleteInstance(ins *types.Instance) error
}
