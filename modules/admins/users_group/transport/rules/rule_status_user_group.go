package rules

import (
	"gocas/modules/admins/common/consts"
	"slices"

	"github.com/go-playground/validator/v10"
)

func RuleStatusUserGroup(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(int)
	return slices.Contains(consts.StatusUserGroupToViewCreate, status)
}

func RuleStatusUserGroupCreate(fl validator.FieldLevel) bool {
	status := fl.Field().Interface().(int)
	return slices.Contains(consts.StatusUserGroupToViewCreate, status)
}
