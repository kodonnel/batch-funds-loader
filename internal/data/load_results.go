package data

// LoadResult defines the structure for an API fund load response
// swagger:model
type LoadResult struct {
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

	// if the load was accepted
	//
	// required: false
	// max length: 10000
	Accepted bool `json:"accepted" validate:"required"`
}

// LoadResults defines a slice of LoadResult
type LoadResults []*LoadResult
