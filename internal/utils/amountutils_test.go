package utils_test

import (
	"testing"

	"github.com/kodonnel/batch-funds-loader/internal/utils"
)

// scenario
// a valid load amount is converted
func TestConvertLoadAmount(t *testing.T) {
	input := "$123.45"
	expected := uint32(12345)
	result, _ := utils.ConvertLoadAmount(input)

	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// an invalid load amount returns an error
func TestConvertLoadAmountWithInvalidAmount(t *testing.T) {
	input := "$12ab3.45"
	_, err := utils.ConvertLoadAmount(input)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}