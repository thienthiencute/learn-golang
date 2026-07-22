package requests

import (
	"context"

	"github.com/go-playground/validator/v10"

	"gocas/modules/admins/common/consts"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users/transport/rules"
)

type UpdateUserReq struct {
	ID        int    `json:"id" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Status    int    `json:"status" validate:"required,ruleStatusUserCreate"`
	StatusStr string `json:"-"`
}

func (req *UpdateUserReq) Validation(ctx context.Context) []*string {
	validation := validator.New()
	validation.RegisterValidation("ruleStatusUserCreate", rules.RuleStatusUserCreate)
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
