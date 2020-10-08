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
// 	- is greater than or equal to 0.0
//  - is less than 1.7976931348623157e+308 (the max for golangs float64 data type)
func ValidateLoadAmount(fl validator.FieldLevel) bool {
	// loadAmount is of the format $123.45
	// ^\$\d+\.\d{2}$
	re := regexp.MustCompile(`^\$\d+\.\d{2}$`)
	strAmount := re.FindAllString(fl.Field().String(), -1)

	if len(strAmount) == 1 {

		fltAmount, err := utils.GetFloatAmount(fl.Field().String())
		if err != nil {
			return false
		}

		if fltAmount < 0.0 {
			return false
		}
		return true
	}
	return false
}

// ValidateID validates that an ID
//	- is greater than 0
//	- less than 2147483647 (the max for golangs int32 data type)
func ValidateID(fl validator.FieldLevel) bool {
	id, err := strconv.ParseInt(fl.Field().String(), 0, 32)
	if err != nil {
		return false
	}
	if id < 1 {
		return false
	}
	return true
}
