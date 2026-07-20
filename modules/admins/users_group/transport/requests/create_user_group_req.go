package requests

import (
	"context"
	"strings"

	"github.com/go-playground/validator/v10"

	"gocas/modules/admins/common/consts"
	"gocas/modules/admins/common/errs"
	"gocas/modules/admins/users_group/transport/rules"
)

type CreateUserGroupReq struct {
	ID 			int 	`json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Status      int    `json:"status"`
}

type UserGroupCreation struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Status      int    `json:"status" validate:"required,ruleStatusUserGroupCreate"`
	StatusStr   string `json:"-"`
}


func (req *UserGroupCreation) Validation(ctx context.Context) []*string {
	validation := validator.New()
	validation.RegisterValidation("ruleStatusUserGroupCreate", rules.RuleStatusUserGroupCreate)
	err := validation.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "Name":
				errName := errs.ErrCreateUserFailed.Error()
				validationErrors = append(validationErrors, &errName)
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

	req.Name = strings.TrimSpace(req.Name)
	req.Description = strings.TrimSpace(req.Description)

	return nil
}
