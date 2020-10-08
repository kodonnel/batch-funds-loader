package data

import (
	"time"

	"github.com/kodonnel/batch-funds-loader/internal/utils"
	"github.com/sirupsen/logrus"
)

// Load defines the structure for an API fund load request
// swagger:model
type Load struct {
	// the id for the load
	//
	// required: true
	ID string `json:"id" validate:"required"`

	// the id for the load customer
	//
	// required: true
	CustomerID string `json:"customer_id" validate:"required"`

	// the amount for this load
	//
	// required: true
	LoadAmount string `json:"load_amount" validate:"required,loadAmount"`

	// the time for this load
	//
	// required: true
	Time time.Time `json:"time" validate:"required"`

	// if the load request was accepted
	//
	// required: false
	Accepted bool `json:"accepted"`
}

// Loads defines a slice of Load
type Loads []*Load

// LoadsDB faux DB for loads
type LoadsDB struct {
	log      *logrus.Logger
	loadList []*Load
}

// NewLoadsDB returns a new faux DB with logger
func NewLoadsDB(l *logrus.Logger) *LoadsDB {

	var loadList = []*Load{}
	return &LoadsDB{l, loadList}
}

// AddLoad adds a new load to the data store
func (dbh *LoadsDB) AddLoad(l Load) {

	dbh.log.Infoln("saving fundsload request", l)
	dbh.loadList = append(dbh.loadList, &l)
}

// IsDuplicate returns true if a load has already been send with the same ID and customerID
func (dbh *LoadsDB) IsDuplicate(load Load) bool {

	dbh.log.Infoln("checking for duplicate fundsload request", load)
	ids := dbh.getLoadIDs(load.CustomerID)
	result := false

	for _, id := range ids {
		if load.ID == id {
			dbh.log.Infoln("found duplicate fundsload request")
			result = true
		}
	}
	return result
}

// GetLoads returns the loads for a customer during the given day
func (dbh *LoadsDB) GetLoads(customer string, accepted bool, start, end time.Time) []*Load {

	dbh.log.Infoln("getting past fundsload requests for customer", customer)
	var loadListForTime = []*Load{}

	for _, load := range dbh.loadList {

		if load.CustomerID == customer && load.Accepted == accepted {
			if utils.IsInTimeSpan(start, end, load.Time) {
				dbh.log.Infoln("found past fundsload request", load.ID)
				loadListForTime = append(loadListForTime, load)
			}
		}
	}

	return loadListForTime
}

// getLoadIDs returns all the load ids for a given customer
func (dbh *LoadsDB) getLoadIDs(customer string) []string {

	dbh.log.Infoln("getting past fundsload request ids for customer", customer)
	var ids []string

	for _, load := range dbh.loadList {
		if load.CustomerID == customer {
			dbh.log.Infoln("found past fundsload request id", load.ID)
			ids = append(ids, load.ID)
		}
	}

	return ids
}
