package validator

import "gopkg.in/go-playground/validator.v9"

var ck = validator.New()

func Check(dst interface{}) error {
	if err := ck.Struct(dst); err != nil {
		return err
	}
	return nil
}
