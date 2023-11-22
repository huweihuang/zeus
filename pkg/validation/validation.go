package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/huweihuang/zeus/pkg/types"
)

func ValidateCreateInstance(ins *types.Instance) error {
	allErrs := field.ErrorList{}
	return allErrs.ToAggregate()
}

func ValidateUpdateInstance(ins *types.Instance) error {
	allErrs := field.ErrorList{}
	return allErrs.ToAggregate()
}

func ValidateDeleteInstance(ins *types.Instance) error {
	allErrs := field.ErrorList{}
	return allErrs.ToAggregate()
}
