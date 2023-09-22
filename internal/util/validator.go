package util

import "github.com/go-playground/validator/v10"

var validate = validator.New(validator.WithRequiredStructEnabled())

func ValidateUrl(url string) bool {
	return validate.Var(url, "required,url") == nil
}
