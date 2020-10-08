package utils

import (
	"strconv"
	"unicode/utf8"
)

// GetFloatAmount trims the dollar sign and returns a float
func GetFloatAmount(amount string) (float64, error) {
	trimmed := trimFirstRune(amount)

	return strconv.ParseFloat(trimmed, 32)
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
