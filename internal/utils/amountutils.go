package utils

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

// utils for working with LoadAmounts

// ConvertLoadAmount returns the uint32 representation of the LoadAmount*100
func ConvertLoadAmount(amount string) (uint32, error) {
	trimmed := trimFirstRune(amount)

	// remove the .
	// -1 means, all occurrences
	noDot := strings.Replace(trimmed, ".", "", -1)

	// use base 10 and 32 bit int
	result, err := strconv.ParseUint(noDot, 10, 32)

	return uint32(result), err
}

func trimFirstRune(s string) string {
	_, i := utf8.DecodeRuneInString(s)
	return s[i:]
}
