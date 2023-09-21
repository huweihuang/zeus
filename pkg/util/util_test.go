package util

import (
	"testing"

	"github.com/huweihuang/gin-api-frame/pkg/constant"
	"github.com/huweihuang/gin-api-frame/pkg/types"
	log "github.com/huweihuang/golib/logger/logrus"
)

func init() {
	log.InitLogger("", "debug", "text", true, false)
}

func TestPrintObjectJson(t *testing.T) {
	ins := &types.Instance{
		InstanceMeta: types.InstanceMeta{
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
	str := PrintObjectJson(ins)
	t.Logf("print json: %s", str)
}

func TestRandStringBytesRmndr(t *testing.T) {
	str := RandStringBytesRmndr(10)
	t.Logf("get random string: %s", str)
}

func TestConvertMapToStr(t *testing.T) {
	testMap := map[string]string{
		"test": "value",
	}
	str := ConvertMapToStr(testMap)
	t.Logf("get string: %s", str)
}

func TestIsInList(t *testing.T) {
	flag := IsInList("test", []string{"test", "test2"})
	t.Logf("get result: %t", flag)
}
