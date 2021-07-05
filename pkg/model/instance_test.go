package model

import (
	"testing"

	"github.com/huweihuang/gin-api-frame/pkg/types"

	"github.com/huweihuang/gin-api-frame/pkg/constant"
	"github.com/huweihuang/gin-api-frame/pkg/util"
)

const (
	insName = "instance"
)

func TestInstanceModel(t *testing.T) {
	TestSetupDB(t)
	t.Run("TestCreateInstance", TestCreateInstance)
	t.Run("TestGetInstanceByHostIDAndAppName", TestGetInstanceByHostIDAndAppName)
	t.Run("TestUpdateInstanceStatus", TestUpdateInstanceStatus)
	t.Run("TestUpdateInstanceImage", TestUpdateInstanceImage)
	t.Run("TestDeleteInstance", TestDeleteInstance)
}

func TestCreateInstance(t *testing.T) {
	TestSetupDB(t)

	e := &types.Instance{
		InstanceMeta: types.InstanceMeta{
			Name:      insName,
			Namespace: "default",
		},
		Spec: types.InstanceSpec{
			HostID:   "xxxxx",
			Image:    "nginx:latest",
			Replicas: 1,
		},
		Status: types.InstanceStatus{
			JobState: constant.JobStateCreating,
			Status:   true,
		},
	}
	if err := CreateInstance(e); err != nil {
		t.Errorf("Failed to create instance: %s", err)
	} else {
		t.Logf("create instance succeed")
	}
}

func TestGetInstanceByHostIDAndAppName(t *testing.T) {
	TestSetupDB(t)
	ins, err := GetInstanceByHostIDAndAppName("xxxx", "app")
	if err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		t.Logf("get instance succeed, %s", util.PrintObjectJson(ins))
	}
}

func TestUpdateInstanceStatus(t *testing.T) {
	TestSetupDB(t)
	err := UpdateInstanceStatus(insName, true)
	if err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		t.Logf("test succeed")
	}
}

func TestUpdateInstanceImage(t *testing.T) {
	TestSetupDB(t)
	err := UpdateInstanceImage(insName, "xxxx")
	if err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		t.Logf("test succeed")
	}
}

func TestDeleteInstance(t *testing.T) {
	TestSetupDB(t)
	err := DeleteInstance(insName)
	if err != nil {
		t.Errorf("test failed: %v", err)
	} else {
		t.Logf("test succeed")
	}
}
