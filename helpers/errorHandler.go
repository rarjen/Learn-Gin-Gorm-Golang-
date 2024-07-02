package helpers

import "github.com/go-playground/validator/v10"

func FormatErrorValidation(err error) []string {
	// Collect all error
	var errors []string

	//Mengubah error validasi menjadi string
	for _, e := range err.(validator.ValidationErrors) {
		errors = append(errors, e.Error())
	}
	return errors
}
