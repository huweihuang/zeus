package model

import (
	"reflect"
	"testing"

	"database/sql/driver"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/huweihuang/golib/db"

	"github.com/huweihuang/zeus/pkg/constant"
	"github.com/huweihuang/zeus/pkg/types"
)

var ins = &types.Instance{
	InstanceMeta: types.InstanceMeta{
		JobID: "xxxxxxxx",
	},
	Spec: types.InstanceSpec{
		Image: "xxxxxxxx",
	},
	Status: types.InstanceStatus{
		JobState: constant.JobStateCreating,
	},
}

func TestInstanceModel(t *testing.T) {
	t.Run("TestCreateInstance", TestCreateInstance)
	t.Run("TestGetInstance", TestGetInstance)
}

func TestCreateInstance(t *testing.T) {
	tests := []struct {
		Name    string
		Query   db.MockQuery
		Args    *types.Instance
		WantErr error
	}{
		{
			Name: "TestCreateInstance",
			Query: db.MockQuery{
				SQL:     "INSERT INTO `instance_tab` (`job_id`,`image`,,`job_state`) VALUES (?,?,?) ON DUPLICATE KEY UPDATE `job_state`=?",
				Args:    []driver.Value{ins.JobID, ins.Spec.Image, ins.Status.JobState, ins.Status.JobState},
				Results: sqlmock.NewResult(1, 1),
			},
			Args:    ins,
			WantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			db.ExpectExec(mock, tt.Query)
			if err := mockDB.CreateInstance(tt.Args); err != tt.WantErr {
				t.Errorf("test failed: %v", err)
				return
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled SQL expectations %v", err)
				return
			}
		})
	}
}

func TestGetInstance(t *testing.T) {
	type args struct {
		hostID string
		name   string
	}
	tests := []struct {
		Name       string
		Query      db.MockQuery
		Args       args
		WantErr    error
		WantResult *types.Instance
	}{
		{
			Name: "TestGetInstance",
			Query: db.MockQuery{
				SQL:  "SELECT * FROM `instance_tab` WHERE host_id=? AND name=? LIMIT 1",
				Args: []driver.Value{ins.Spec.HostID, ins.Name},
				Rows: mock.
					NewRows([]string{`job_id`, `image`, `job_state`}).
					AddRow(ins.JobID, ins.Spec.Image, ins.Status.JobState),
			},
			Args: args{
				hostID: ins.Spec.HostID,
				name:   ins.Name,
			},
			WantErr:    nil,
			WantResult: ins,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			db.ExpectQuery(mock, tt.Query)
			got, err := mockDB.GetInstance(tt.Args.hostID, tt.Args.name)
			if err != tt.WantErr {
				t.Errorf("test failed: %v", err)
				return
			}
			if !reflect.DeepEqual(got, tt.WantResult) {
				t.Errorf("got = %v, want %v", got, tt.WantResult)
			}
			if err := mock.ExpectationsWereMet(); err != nil {
				t.Errorf("unfulfilled SQL expectations %v", err)
				return
			}
		})
	}
}
