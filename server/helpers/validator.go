package helpers

import (
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type ErrorResponse struct {
	FailedField string `json:"failedField"`
	Tag         string `json:"tag"`
	Value       string `json:"value"`
}

var validate = validator.New()

func ValidateStruct(c *fiber.Ctx, data interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(data)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var param string
			splitted := strings.Split(err.StructNamespace(), ".")
			field := splitted[len(splitted)-1]
			if err.Tag() == "email" {
				param = "Must Email"
			} else if err.Tag() == "required" {
				param = field + " " + "required"
			} else {
				param = err.Param()
			}
			element := ErrorResponse{
				FailedField: err.StructNamespace(),
				Tag:         err.Tag(),
				Value:       param,
			}
			errors = append(errors, &element)
		}
	}
	return errors
}
