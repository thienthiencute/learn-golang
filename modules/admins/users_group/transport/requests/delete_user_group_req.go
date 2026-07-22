package requests

import (
	"context"

	"github.com/go-playground/validator/v10"

	"gocas/modules/admins/common/errs"
)

type UserGroupDeleteReq struct {
	Ids []string `json:"ids" validate:"required"`
}

func (req *UserGroupDeleteReq) Validation(ctx context.Context) []*string {
	validate := validator.New()
	err := validate.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "Ids":
				errId := errs.ErrIDUserValidate.Error()
				validationErrors = append(validationErrors, &errId)
			}
		}
	}
	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
