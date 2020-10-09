
package handlers_test

import (
	"testing"
	"time"
	"reflect"

	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/go-playground/validator"
	"github.com/sirupsen/logrus"
	"github.com/kodonnel/batch-funds-loader/internal/handlers"

)

// scenario
// test invalid input results in error
func TestProcessLoadRequestError(t *testing.T) {

			logger := &logrus.Logger{}

			// create db instance
			db := data.NewLoadsDB(logger)

			// create validator
			v := validator.New()
			v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
			v.RegisterValidation("identifier", data.ValidateID)

			// req handlers
			loadsHandler := handlers.NewLoads(logger, db, v)

			load := new(data.Load)
			load.CustomerID = "123"
			load.ID = "1234"
			load.Time = time.Now()
			load.LoadAmount = "5000.45"

			_, err := loadsHandler.ProcessLoadRequest(*load)

			if err == nil {
				t.Errorf("expected validation errors, received none")
			}
		}

// scenario
// test valid input does not have error
func TestProcessLoadRequestAccepted(t *testing.T) {

	logger := &logrus.Logger{}

	// create db instance
	db := data.NewLoadsDB(logger)

	// create validator
	v := validator.New()
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
	v.RegisterValidation("identifier", data.ValidateID)

	// req handlers
	loadsHandler := handlers.NewLoads(logger, db, v)

	loc, _ := time.LoadLocation("America/Toronto")

	loadtime := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)

	load := new(data.Load)
	load.CustomerID = "123"
	load.ID = "1234"
	load.Time = loadtime
	load.LoadAmount = "$5000.00"


	expected := new(data.Load)
	expected.CustomerID = "123"
	expected.ID = "1234"
	expected.Accepted = true	
	
	result, _ := loadsHandler.ProcessLoadRequest(*load)

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// test valid input does not have error
// load is not accepted due to max daily amount exceeded
func TestProcessLoadRequestDeclinedMaxDailyAmount(t *testing.T) {

	logger := &logrus.Logger{}

	// create db instance
	db := data.NewLoadsDB(logger)

	// create validator
	v := validator.New()
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
	v.RegisterValidation("identifier", data.ValidateID)

	// req handlers
	loadsHandler := handlers.NewLoads(logger, db, v)

	loc, _ := time.LoadLocation("America/Toronto")

	loadtime := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)

	load := new(data.Load)
	load.CustomerID = "123"
	load.ID = "1234"
	load.Time = loadtime
	load.LoadAmount = "$5000.01"


	expected := new(data.Load)
	expected.CustomerID = "123"
	expected.ID = "1234"
	expected.Accepted = false	
	
	result, _ := loadsHandler.ProcessLoadRequest(*load)

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// test valid input does not have error
// load is not accepted due to max daily number of loads exceeded
func TestProcessLoadRequestDeclinedMaxDailyNumber(t *testing.T) {

	logger := &logrus.Logger{}

	// create db instance
	db := data.NewLoadsDB(logger)

	// add pre-existing loads
	loc, _ := time.LoadLocation("America/Toronto")

	loadtime := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = loadtime
	load.LoadAmount = "$50.00"
	load.Accepted = true
	db.AddLoad(*load)

	loadtime = time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	load2 := new(data.Load)
	load2.CustomerID = "1234"
	load2.ID = "1235"
	load2.Time = loadtime
	load2.LoadAmount = "$50.00"
	load2.Accepted = true
	db.AddLoad(*load2)

	loadtime = time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	load3 := new(data.Load)
	load3.CustomerID = "1234"
	load3.ID = "1236"
	load3.Time = loadtime
	load3.LoadAmount = "$50.00"
	load3.Accepted = true
	db.AddLoad(*load3)

	// create validator
	v := validator.New()
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
	v.RegisterValidation("identifier", data.ValidateID)

	// req handlers
	loadsHandler := handlers.NewLoads(logger, db, v)

	loadtime = time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	given := new(data.Load)
	given.CustomerID = "1234"
	given.ID = "1238"
	given.Time = loadtime
	given.LoadAmount = "$0.01"

	expected := new(data.Load)
	expected.CustomerID = "1234"
	expected.ID = "1238"
	expected.Accepted = false	
	
	result, _ := loadsHandler.ProcessLoadRequest(*given)

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// test valid input does not have error
// load is not accepted due to max weekly amount exceeded
func TestProcessLoadRequestDeclinedMaxWeeklyAmount(t *testing.T) {

	logger := &logrus.Logger{}

	// create db instance
	db := data.NewLoadsDB(logger)

	// add pre-existing loads
	loc, _ := time.LoadLocation("America/Toronto")

	loadtime := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"
	load.Time = loadtime
	load.LoadAmount = "$5000.00"
	load.Accepted = true
	db.AddLoad(*load)

	loadtime = time.Date(2020, 9, 28, 20, 0, 0, 0, loc)

	load2 := new(data.Load)
	load2.CustomerID = "1234"
	load2.ID = "1235"
	load2.Time = loadtime
	load2.LoadAmount = "$5000.00"
	load2.Accepted = true
	db.AddLoad(*load2)

	loadtime = time.Date(2020, 9, 29, 20, 0, 0, 0, loc)

	load3 := new(data.Load)
	load3.CustomerID = "1234"
	load3.ID = "1236"
	load3.Time = loadtime
	load3.LoadAmount = "$5000.00"
	load3.Accepted = true
	db.AddLoad(*load3)

	loadtime = time.Date(2020, 9, 30, 20, 0, 0, 0, loc)

	load4 := new(data.Load)
	load4.CustomerID = "1234"
	load4.ID = "1237"
	load4.Time = loadtime
	load4.LoadAmount = "$5000.00"
	load4.Accepted = true
	db.AddLoad(*load4)

	// create validator
	v := validator.New()
	v.RegisterValidation("loadAmount", data.ValidateLoadAmount)
	v.RegisterValidation("identifier", data.ValidateID)

	// req handlers
	loadsHandler := handlers.NewLoads(logger, db, v)

	loadtime = time.Date(2020, 10, 1, 20, 0, 0, 0, loc)

	given := new(data.Load)
	given.CustomerID = "1234"
	given.ID = "1238"
	given.Time = loadtime
	given.LoadAmount = "$0.01"

	expected := new(data.Load)
	expected.CustomerID = "1234"
	expected.ID = "1238"
	expected.Accepted = false	
	
	result, _ := loadsHandler.ProcessLoadRequest(*given)

	if !reflect.DeepEqual(result,expected) {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}