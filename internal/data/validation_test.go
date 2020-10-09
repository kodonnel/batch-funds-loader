package data_test

import (
	"testing"
	"time"

	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/go-playground/validator"

)

// scenario
// no validation errors
func TestValidateID(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "$123.45"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err != nil {
    	t.Errorf("expected no validation errors, received %s", err)
  	}
}

// scenario
// load amount is missing $
func TestValidateIDWithLoadAmountMissingDollar(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "123.45"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}

// scenario
// load amount is missing .
func TestValidateIDWithLoadAmountMissingDot(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "$12345"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}

// scenario
// load amount is too big
func TestValidateIDWithLoadAmountTooBig(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "$99999999999999.99"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}

// scenario
// id is too small
func TestValidateIDWithSmallID(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "-4"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "$123.45"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}

// scenario
// id is too big
func TestValidateIDWithBigID(t *testing.T) {

	load := new(data.Load)
	load.CustomerID = "9999999999999999"
	load.ID = "1234"
	load.Time = time.Now()
	load.LoadAmount = "$123.45"

	v := validator.New()
	v.RegisterValidation("identifier", data.ValidateID)
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)

	err := v.Struct(load)

	if err == nil {
    	t.Errorf("expected validation errors, none received")
  	}
}