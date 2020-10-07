package data

import (
	"fmt"
	"time"

	"github.com/kodonnel/batch-funds-loader/internal/utils"
)

// Load defines the structure for an API fund load request
// swagger:model
type Load struct {
	// the id for the load
	//
	// required: false
	// min:1
	ID string `json:"id" validate:"required"`

	// the id for the load customer
	//
	// required: true
	// max length: 255
	CustomerID string `json:"customer_id" validate:"required"`

	// the amount for this load
	//
	// required: false
	// max length: 10000
	LoadAmount string `json:"load_amount" validate:"required"`

	// the time for this load
	//
	// required: true
	// min: 0.01
	Time time.Time `json:"time" validate:"required"`

	// if the load request was accepted
	Accepted bool `json:"accepted"`
}

// Loads defines a slice of Load
type Loads []*Load

// AddLoad adds a new load to the data store
func AddLoad(l Load) {

	loadList = append(loadList, &l)
}

// IsDuplicate returns true if a load has already been send with the same ID and customerID
func IsDuplicate(load Load) bool {

	ids := GetLoadIDs(load.CustomerID)
	result := false

	for _, id := range ids {

		if load.ID == id {
			result = true
		}
	}
	return result
}

// GetLoads returns the loads for a customer during the given day
func GetLoads(customer string, accepted bool, start, end time.Time) []*Load {

	var loadListForTime = []*Load{}
	fmt.Printf("checking loads for customer %s \n", customer)

	for _, load := range loadList {
		fmt.Printf("check if load is in span: %s %s %s %s \n", load.CustomerID, load.ID, load.LoadAmount, load.Time)

		if load.CustomerID == customer && load.Accepted {
			fmt.Printf("check if load is in span: %s %s %s %s \n", load.CustomerID, load.ID, load.LoadAmount, load.Time)
			if utils.InTimeSpan(start, end, load.Time) {
				fmt.Printf("it was \n")

				loadListForTime = append(loadListForTime, load)
			}
		}
	}

	return loadListForTime
}

// GetLoadIDs returns all the load ids for a customer
func GetLoadIDs(customer string) []string {
	var ids []string

	for _, load := range loadList {

		if load.CustomerID == customer {
			ids = append(ids, load.ID)
		}
	}

	return ids
}

// hardcoded list instead of database
var loadList = []*Load{}
