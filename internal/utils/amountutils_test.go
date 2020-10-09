package utils_test

import (
	"testing"

	"github.com/kodonnel/batch-funds-loader/internal/utils"
)

func TestConvertLoadAmount(t *testing.T) {
	input := "$123.45"
	expected := uint32(12345)
	result, _ := utils.ConvertLoadAmount(input)

	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}
