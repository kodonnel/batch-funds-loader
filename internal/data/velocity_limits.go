package data

import "fmt"

// ErrVelocityLimitNotFound is an error raised when a product can not be found in the data store
var ErrVelocityLimitNotFound = fmt.Errorf("Velocity Limit not found")

// VelocityLimit defines the structure for an customers limits
// swagger:model
type VelocityLimit struct {
	// previous load ids for the customer
	//
	// required: false
	// min:1
	LoadIDs []string

	// the id for the load
	//
	// required: false
	// min:1
	CustomerID string

	// the amount for this load
	//
	// required: false
	// max length: 10000
	DailyAmount float32

	// the weekly amount for this load
	//
	// required: false
	// max length: 10000
	WeeklyAmount float32

	// the amount for this load
	//
	// required: false
	// max length: 10000
	DailyLoads int
}

// VelocityLimits defines a slice of VelocityLimit
var VelocityLimits []*VelocityLimit

// AddVelocityLimit adds a new velocity limit to the data store
func AddVelocityLimit(vl VelocityLimit) {

	limitList = append(limitList, &vl)
}

// CheckLoadIDExistsForCustomer check if this is a duplicate load
func CheckLoadIDExistsForCustomer(id string, customer string) bool {

	fmt.Printf("Checking existance for load %s for customer %s \n", id, customer)

	idx := findIndexByCustomerID(customer)

	fmt.Printf("customer index was %d \n", idx)

	if idx == -1 {
		// no customer exists yet
		return false
	}

	cust := limitList[idx]

	fmt.Printf("customer info is %s \n", cust.CustomerID)
	fmt.Println(cust.LoadIDs)

	for _, existingid := range cust.LoadIDs {
		fmt.Printf("customer already has id: %s \n", existingid)
		if id == existingid {
			fmt.Println("id was found")
			return true
		}
	}

	return false
}

// findIndex finds the index of a customer in the data store
// returns -1 when no customer can be found
func findIndexByCustomerID(id string) int {
	for i, p := range limitList {
		if p.CustomerID == id {
			return i
		}
	}

	return -1
}

// hardcoded list instead of database
var limitList = []*VelocityLimit{}
