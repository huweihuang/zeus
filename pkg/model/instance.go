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

func (m *DBMng) CreateInstance(ins *types.Instance) error {
	err := m.db.Clauses(clause.OnConflict{
		DoUpdates: clause.Assignments(map[string]interface{}{"replicas": ins.Spec.Replicas,
			"status": ins.Status.Status, "job_state": ins.Status.JobState})}).Create(ins).Error
	if err != nil {
		return err
	}
	log.Logger.WithField("[instance]", util.PrintObjectJson(ins)).Debug("create instance in db")
	return nil
}

func (m *DBMng) GetInstance(hostID, name string) (*types.Instance, error) {
	ins := &types.Instance{}
	err := m.db.Where("host_id= ? AND name=?", hostID, name).Take(ins).Error
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

func (m *DBMng) UpdateInstanceStatus(name string, status bool) error {
	err := m.db.Where("name = ?", name).Update("status", status).Error
	if err != nil {
		return fmt.Errorf("failed to update instance status by name, err: %v", err)
	}

	log.Logger.WithFields(logrus.Fields{
		"insName": name, "status": status,
	}).Info("update instance status in db")
	return nil
}

func (m *DBMng) UpdateInstanceImage(name, image string) error {
	err := m.db.Where("name = ?", name).Update("image", image).Error
	if err != nil {
		return fmt.Errorf("failed to update instance image by job_id, err: %v", err)
	}
	log.Logger.WithFields(logrus.Fields{"insName": name, "image": image}).Info("update job image in db")
	return nil
}

func (m *DBMng) DeleteInstance(insName string) error {
	err := m.db.Where("name = ?", insName).Delete(types.Instance{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete instance in db, err: %v", err)
	}
	log.Logger.WithField("[instance]", insName).Debug("delete instance from db")
	return nil
}
