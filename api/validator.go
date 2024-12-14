package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/imran4u/simple-bank/util"
)

var validCurrency validator.Func = func(fieldLabel validator.FieldLevel) bool {
	if currency, ok := fieldLabel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}

	return false
}
