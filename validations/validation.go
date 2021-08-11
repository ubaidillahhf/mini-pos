package validations

import "gopkg.in/go-playground/validator.v9"

func Validate(context interface{}) error {
	v := validator.New()
	errValidation := v.Struct(context)
	if errValidation != nil {
		return errValidation
	}
	return nil
}
