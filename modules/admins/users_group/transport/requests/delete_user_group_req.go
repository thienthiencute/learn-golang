package requests

import (
	"gocas/modules/admins/common/errs"

	"github.com/go-playground/validator/v10"
)

type UserGroupDeleteReq struct {
	Ids []string `json:"ids" validate:"required"`
}

func (req *UserGroupDeleteReq) Validation() error {
	validate := validator.New()
	err := validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Ids":
				return errs.ErrIDUserValidate
			}
		}
	}
	return nil
}
