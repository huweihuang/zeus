package validation

import (
	"k8s.io/apimachinery/pkg/util/validation/field"

	"github.com/huweihuang/zeus/pkg/types"
)

// 校验创建请求
func ValidateCreateInstance(ins *types.Instance) field.ErrorList {
	allErrs := field.ErrorList{}
	return allErrs
}

// 校验更新请求
func ValidateUpdateInstance(ins *types.Instance) field.ErrorList {
	allErrs := field.ErrorList{}
	return allErrs
}

// 校验删除请求
func ValidateDeleteInstance(ins *types.Instance) field.ErrorList {
	allErrs := field.ErrorList{}
	return allErrs
}
