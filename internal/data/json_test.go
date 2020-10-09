package data_test

import (
	"testing"

	"github.com/kodonnel/batch-funds-loader/internal/data"
)

// scenario
// adding the same funds load twice
func TestMarshalJSON(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Accepted = true
	
	_, err := data.MarshalJSON(*load)

	if err != nil {
		t.Errorf("failed unable to marshal json, got %v", err)
	}
}