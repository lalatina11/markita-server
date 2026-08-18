package validator

// utils/validator.go

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
	_validator "github.com/go-playground/validator/v10"
)

var validate = _validator.New()

func init() {
	// Use json tag name in error messages instead of struct field name
	// So you get "email" not "Email" in errors
	validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

type ValidationErrorDetail struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

func Validate(payload interface{}) []ValidationErrorDetail {
	err := validate.Struct(payload)
	if err == nil {
		return nil
	}

	var errs []ValidationErrorDetail
	for _, e := range err.(validator.ValidationErrors) {
		errs = append(errs, ValidationErrorDetail{
			Field:   e.Field(),
			Message: buildMessage(e),
		})
	}
	return errs
}

func buildMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return fmt.Sprintf("%s is required", e.Field())
	case "email":
		return fmt.Sprintf("%s must be a valid email", e.Field())
	case "min":
		return fmt.Sprintf("%s must be at least %s characters", e.Field(), e.Param())
	case "max":
		return fmt.Sprintf("%s must be at most %s characters", e.Field(), e.Param())
	case "oneof":
		return fmt.Sprintf("%s must be one of: %s", e.Field(), e.Param())
	default:
		return fmt.Sprintf("%s is invalid", e.Field())
	}
}
