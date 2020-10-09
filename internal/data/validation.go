package data

import (
	"regexp"
	"strconv"

	"github.com/go-playground/validator"
	"github.com/kodonnel/batch-funds-loader/internal/utils"
)

// ValidateLoadAmount validates that the load amount
//	- starts with a $
//	- followed by 1 or more digits
//  - followed by a .
//  - followed by 2 digits
// 	- is greater than or equal to 0
//  - is less than $42,949,672.95 (the max for golangs uint32 data type)
func ValidateLoadAmount(fl validator.FieldLevel) bool {
	// loadAmount is of the format $123.45
	// ^\$\d+\.\d{2}$
	re := regexp.MustCompile(`^\$\d+\.\d{2}$`)
	strAmount := re.FindAllString(fl.Field().String(), -1)

	if len(strAmount) == 1 {

		_, err := utils.ConvertLoadAmount(fl.Field().String())
		if err != nil {
			return false
		}

		return true
	}
	return false
}

// ValidateID validates that an ID
//	- is greater than 0
//	- less than 4294967295 (the max for golangs uint32 data type)
func ValidateID(fl validator.FieldLevel) bool {

	// use base 10 and 32 bit unsigned int
	_, err := strconv.ParseUint(fl.Field().String(), 10, 32)
	if err != nil {
		return false
	}

	return true
}
