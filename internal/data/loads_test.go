package data_test

import (
	"testing"
	"time"

	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/sirupsen/logrus"
)

// scenario
// adding the same funds load twice
func TestIsDuplicateWithExistingLoad(t *testing.T) {

	logger := &logrus.Logger{}

	db := data.NewLoadsDB(logger)

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"

	db.AddLoad(*load)
	
	result := db.IsDuplicate(*load)
	expected := true

	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// adding two different funds loads
// different customerIDs
func TestIsDuplicateWithNewCustomerId(t *testing.T) {

	logger := &logrus.Logger{}

	db := data.NewLoadsDB(logger)

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"

	db.AddLoad(*load)
	
	nload := new(data.Load)
	nload.CustomerID = "123"
	nload.ID = "1234"
	result := db.IsDuplicate(*nload)
	expected := false

	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// adding two different funds loads
// different loadIDs
func TestIsDuplicateWithNewLoadId(t *testing.T) {

	logger := &logrus.Logger{}

	db := data.NewLoadsDB(logger)

	load := new(data.Load)
	load.CustomerID = "1234"
	load.ID = "1234"

	db.AddLoad(*load)
	
	nload := new(data.Load)
	nload.CustomerID = "1234"
	nload.ID = "123"
	result := db.IsDuplicate(*nload)
	expected := false

	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// add three loads - all in a time range
func TestGetLoads(t *testing.T) {

	logger := &logrus.Logger{}

	db := data.NewLoadsDB(logger)

	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	time1 := time.Date(2020, 10, 7, 19, 59, 59, 999999999, loc)
	time2 := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)
	time3 := time.Date(2020, 10, 7, 20, 0, 0, 1, loc)


	load := new(data.Load)
	load.CustomerID = "1234"
	load.Accepted = true
	load.ID = "1234"
	load.Time = time2

	db.AddLoad(*load)

	load.ID = "123"
	load.Time = time1
	db.AddLoad(*load)

	load.ID = "1"
	load.Time = time3
	db.AddLoad(*load)

	result := db.GetLoads("1234", true, time1, time3)
	expected := 3

	if len(result) != expected {
		t.Errorf("failed expected %v got %v", expected, len(result))
	}
}

// scenario
// add three loads to the database
// test getting a time range that only includes 1
func TestGetLoadsWithOutOfRange(t *testing.T) {

	logger := &logrus.Logger{}

	db := data.NewLoadsDB(logger)

	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	time1 := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)
	beforeTime := time.Date(2020, 10, 7, 19, 59, 58, 0, loc)
	afterTime := time.Date(2020, 10, 9, 20, 0, 1, 0, loc)


	load := new(data.Load)
	load.CustomerID = "1234"
	load.Accepted = true
	load.ID = "1234"
	load.Time = time1

	db.AddLoad(*load)

	load.ID = "123"
	load.Time = beforeTime
	db.AddLoad(*load)

	load.ID = "1"
	load.Time = afterTime
	db.AddLoad(*load)

	result := db.GetLoads("1234", true, time1, time1)
	expected := 1

	if len(result) != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}