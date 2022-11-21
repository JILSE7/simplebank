package api

import (
	"github.com/JILSE7/simplebank/utils"
	"github.com/go-playground/validator/v10"
)

var validCurrency validator.Func = func(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		//check if currency is supported
		return utils.IsValidCurrency(currency)
	}

	return false
}
