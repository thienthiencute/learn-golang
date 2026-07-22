package requests

import (
	"context"

	"github.com/go-playground/validator/v10"

	"gocas/modules/admins/common/consts"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/transport/rules"
)

type UserGroupUpdateStatusReq struct {
	ID        int    `json:"id" validate:"required"`
	Status    int    `json:"status" validate:"required,ruleStatusUserGroupCreate"`
	StatusStr string `json:"-"`
}

func (req *UserGroupUpdateStatusReq) Validation(ctx context.Context) []*string {
	validation := validator.New()
	validation.RegisterValidation("ruleStatusUserGroupCreate", rules.RuleStatusUserGroupCreate)
	err := validation.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "Status":
				errStatus := errs.ErrStatusNotFound.Error()
				validationErrors = append(validationErrors, &errStatus)
			}
		}
	}
	if validationErrors != nil {
		return validationErrors
	}

	req.StatusStr = consts.MapStatusIntToString[req.Status]

	return nil
}
