package validation

import "github.com/go-playground/validator/v10"

type ErrorResponse struct {
	FailedField string `json:"failedField,omitempty"`
	Tag         string `json:"tag,omitempty"`
	Value       string `json:"value,omitempty"`
}

var validate = validator.New()

func ValidateStruct(i interface{}) []*ErrorResponse {
	var errors []*ErrorResponse
	err := validate.Struct(i)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
