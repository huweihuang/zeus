package model

import (
	"errors"
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"

	log "github.com/huweihuang/golib/logger/logrus"
	errConst "github.com/huweihuang/zeus/pkg/errors"
	"github.com/huweihuang/zeus/pkg/types"
	"github.com/huweihuang/zeus/pkg/util"
)

const (
	tableInstance = "t_instance"
)

// 创建数据库instance
func CreateInstance(ins *types.Instance) error {
	err := DB.Table(tableInstance).Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{"replicas": ins.Spec.Replicas,
			"status": ins.Status.Status, "job_state": ins.Status.JobState})}).Create(ins).Error
	if err != nil {
		return err
	}
	log.Logger.WithField("[instance]", util.PrintObjectJson(ins)).Debug("create instance in db")
	return nil
}

// 查询数据库instance
func GetInstanceByHostIDAndAppName(hostID, name string) (*types.Instance, error) {
	ins := &types.Instance{}
	err := DB.Table(tableInstance).Where("host_id= ? AND name=?", hostID, name).Take(ins).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errConst.ErrInstanceNotFound
		}
		return nil, err
	}

	log.Logger.WithField(
		"[instance]", util.PrintObjectJson(ins),
	).Debug("get instance from db by host id")
	return ins, nil
}

// 更新数据库instance状态
func UpdateInstanceStatus(name string, status bool) error {
	err := DB.Table(tableInstance).Where("name = ?", name).Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update instance status by name, err: %v", err)
	}

	log.Logger.WithFields(logrus.Fields{
		"insName": name, "status": status,
	}).Info("update instance status in db")
	return nil
}

// 更新数据库instance镜像
func UpdateInstanceImage(name, image string) error {
	err := DB.Table(tableInstance).Where("name = ?", name).Update("image", image).Error
	if err != nil {
		return fmt.Errorf("failed to update instance image by job_id, err: %v", err)
	}
	log.Logger.WithFields(logrus.Fields{"insName": name, "image": image}).Info("update job image in db")
	return nil
}

// 删除数据库instance
func DeleteInstance(insName string) error {
	err := DB.Table(tableInstance).Where("name = ?", insName).Delete(types.Instance{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete instance in db, err: %v", err)
	}
	log.Logger.WithField("[instance]", insName).Debug("delete instance from db")
	return nil
}
