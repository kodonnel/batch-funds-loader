package data

import (
	"regexp"

	"github.com/go-playground/validator"
)

// ValidateLoadAmount validates that the load amount
//	- starts with a $
//	- followed by 1 or more digits
//  - followed by a .
//  - followed by 2 digits
func ValidateLoadAmount(fl validator.FieldLevel) bool {
	// loadAmount is of the format $123.45
	// ^\$\d+\.\d{2}$
	re := regexp.MustCompile(`^\$\d+\.\d{2}$`)
	loadAmount := re.FindAllString(fl.Field().String(), -1)

	if len(loadAmount) == 1 {
		return true
	}
	return false
}
