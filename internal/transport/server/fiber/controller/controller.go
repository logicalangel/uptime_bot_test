package controller

import (
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/logicalangel/tashil_test/internal/consts"
	"github.com/logicalangel/tashil_test/internal/transport/server/fiber/dto"
)

var Validator = validator.New()

func ValidateBody(body interface{}) error {
	err := Validator.Struct(body)
	if err != nil {
		errResponse := dto.ResponseError{
			Message: consts.ErrBadBodyRequest.Error(),
		}
		var fieldErrors validator.ValidationErrors
		if errors.As(err, &fieldErrors) {
			for _, err := range fieldErrors {
				errResponse.Fields = append(errResponse.Fields, dto.FieldError{
					Field: err.Field(),
					Tag:   err.Tag(),
					Value: err.Param(),
				})
			}
		}
		return errResponse
	}

	return nil
}
