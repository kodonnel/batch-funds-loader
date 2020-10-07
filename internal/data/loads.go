package data

import (
	"time"
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
}

// Loads defines a slice of Load
type Loads []*Load

// AddLoad adds a new load to the data store
func AddLoad(l Load) {

}
