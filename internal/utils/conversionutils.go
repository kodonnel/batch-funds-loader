package utils

import (
	"fmt"
	"strconv"
	"unicode/utf8"
)

// GetFloatAmount trims the dollar sign and returns a float
func GetFloatAmount(amount string) float64 {
	trimmed := trimFirstRune(amount)

	f, err := strconv.ParseFloat(trimmed, 64)
	if err != nil {
		fmt.Println("Bad amount")
	}

	return f
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
