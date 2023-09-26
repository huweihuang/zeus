package validation

import (
	"testing"

	"github.com/huweihuang/zeus/pkg/types"
)

var (
	ins = &types.Instance{}
)

func TestValidateCreateInstance(t *testing.T) {
	errs := ValidateCreateInstance(ins)
	if len(errs) != 0 {
		t.Errorf("test failed, err: %v", errs)
	} else {
		t.Logf("validate succeed")
	}
}

func TestValidateUpdateInstance(t *testing.T) {
	errs := ValidateUpdateInstance(ins)
	if len(errs) != 0 {
		t.Errorf("test failed, err: %v", errs)
	} else {
		t.Logf("validate succeed")
	}
}

func TestValidateDeleteInstance(t *testing.T) {
	errs := ValidateDeleteInstance(ins)
	if len(errs) != 0 {
		t.Errorf("test failed, err: %v", errs)
	} else {
		t.Logf("validate succeed")
	}
}
