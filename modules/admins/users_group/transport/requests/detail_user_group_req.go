package requests

import (
	"context"

	"github.com/go-playground/validator/v10"
)

type UserGroupDetail struct {
	ID int64 `json:"id" validate:"required"`
}

func (req *UserGroupDetail) Validation(ctx context.Context) []*string {
	validation := validator.New()
	err := validation.Struct(req)
	var validationErrors []*string
	if err != nil {
		for _, vErr := range err.(validator.ValidationErrors) {
			switch vErr.Field() {
			case "ID":
				errMsg := "ID is required"
				validationErrors = append(validationErrors, &errMsg)
			}
		}
	}
	if validationErrors != nil {
		return validationErrors
	}
	return nil
}
