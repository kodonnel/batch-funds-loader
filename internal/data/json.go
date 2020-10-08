package data

import (
	"encoding/json"
)

// MarshalJSON creates a temporarty
func MarshalJSON(load Load) ([]byte, error) {
	var tmp struct {
		// the id for the load
		//
		// required: true
		ID string `json:"id" validate:"required,identifier"`

		// the id for the load customer
		//
		// required: true
		CustomerID string `json:"customer_id" validate:"required,identifier"`

		// if the load request was accepted
		//
		// required: false
		Accepted bool `json:"accepted"`
	}
	tmp.ID = load.ID
	tmp.CustomerID = load.CustomerID
	tmp.Accepted = load.Accepted
	return json.Marshal(&tmp)
}
