package controller_test

import (
	"testing"

	"github.com/huweihuang/gin-api-frame/pkg/types"
)

var (
	ins = &types.Instance{}
)

func TestInstanceController_CreateInstance(t *testing.T) {
	t.Run("CreateInstance", func(t *testing.T) {
		if err := InsCtrl.CreateInstance(ins); err != nil {
			t.Errorf("InstanceController.CreateInstance() error = %v", err)
		}
	})
}

func TestInstanceController_GetInstance(t *testing.T) {
	t.Run("GetInstance", func(t *testing.T) {
		got, err := InsCtrl.GetInstance(ins)
		if err != nil {
			t.Errorf("InstanceController.GetInstance() error = %v", err)
			return
		}
		t.Logf("InstanceController.GetInstance() = %v", got)
	})
}

func TestInstanceController_UpdateInstance(t *testing.T) {
	t.Run("UpdateInstance", func(t *testing.T) {
		if err := InsCtrl.UpdateInstance(ins); err != nil {
			t.Errorf("InstanceController.UpdateInstance() error = %v", err)
		}
	})
}

func TestInstanceController_DeleteInstance(t *testing.T) {
	t.Run("DeleteInstance", func(t *testing.T) {
		if err := InsCtrl.DeleteInstance(ins); err != nil {
			t.Errorf("InstanceController.DeleteInstance() error = %v", err)
		}
	})
}
