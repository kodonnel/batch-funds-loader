package handlers

import (
	"fmt"

	"github.com/go-playground/validator"
	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/kodonnel/batch-funds-loader/internal/utils"
	"github.com/sirupsen/logrus"
)

// ErrDuplicateFound indicates that a load with the same id and customer_id was already processed
var ErrDuplicateFound = fmt.Errorf("Duplicate Load Request Found")

// ErrValidationError indicates that a load did not have valid fields
var ErrValidationError = fmt.Errorf("Unable to validate Load Request")

const maxLoadsPerDay = 3

// dollar amounts multiplied by 100 (500 = $5)
const maxLoadAmountPerDay = 500000
const maxLoadAmountPerWeek = 2000000

// Loads handler for getting and updating funds load requests
type Loads struct {
	l  *logrus.Logger
	db *data.LoadsDB
	v  *validator.Validate
}

// NewLoads returns a new products handler with the given logger
func NewLoads(l *logrus.Logger, db *data.LoadsDB, v *validator.Validate) *Loads {
	return &Loads{l, db, v}
}

// ProcessLoadRequest processes the load request and return the result
// (apply business logic)
func (lh *Loads) ProcessLoadRequest(req data.Load) (*data.Load, error) {

	lh.l.Infoln("processing load request", req)

	// validate the loadFunds request
	// if any fields are invalid, skip
	err := lh.v.Struct(req)

	if err != nil {
		lh.l.Errorln("Unable to validate", err)
		lh.l.Infoln("Could not validate load request, skipping", req)

		for _, e := range err.(validator.ValidationErrors) {
			lh.l.Errorln("Validate error", e)
		}
		return nil, ErrValidationError
	}

	var loadResult *data.Load

	// check if duplicate
	if !lh.db.IsDuplicate(req) {

		req.Accepted = (lh.isWithinDailyLimits(req) && lh.isWithinWeeklyLimits(req))

		lh.db.AddLoad(req)

		loadResult = new(data.Load)
		loadResult.CustomerID = req.CustomerID
		loadResult.ID = req.ID
		loadResult.Accepted = req.Accepted
		return loadResult, nil
	}

	return nil, ErrDuplicateFound
}

// isWithingDailyLimits checks of a given load would be within the daily limits
// including
//    * A maximum of 3 loads can be performed per day, regardless of amount
//	     - loads that were not accepted do not count against the maximum
//    * A maximum of $5,000 can be loaded per day
//		 - loads that were not accepted do not count against the maximum
func (lh *Loads) isWithinDailyLimits(load data.Load) bool {
	lh.l.Infoln("checking daily limits for load", load)

	result := true

	// get all the loads for the same day as the requested load
	loads := lh.db.GetLoads(load.CustomerID, true, utils.GetStartOfDay(load.Time), utils.GetEndOfDay(load.Time))

	if len(loads) >= maxLoadsPerDay {
		lh.l.Infoln("exceeded maximum loads per day")
		result = false
	}

	dailySum, err := lh.addLoadAmounts(loads)

	if err != nil {
		// defensive coding
		// input validation should prevent this
		// if we could not calculate the daily sum, do not accept the load
		lh.l.Errorln("unable to calculate daily sum", err)
		return false
	}

	amount, err := utils.ConvertLoadAmount(load.LoadAmount)

	if err != nil {
		// defensive coding
		// input validation should prevent this
		lh.l.Errorln("unable to get amount from fundsload request", err)
		return false
	}

	if (dailySum + amount) > maxLoadAmountPerDay {
		lh.l.Infoln("exceeded maximum load amount per day")
		result = false
	}

	return result
}

// isWithingWeeklyLimits checks of a given load would be within the weekly limits
// including
//    * A maximum of $20,000 can be loaded per day
//		 - loads that were not accepted do not count against the maximum
func (lh *Loads) isWithinWeeklyLimits(load data.Load) bool {
	lh.l.Infoln("checking weekly limits for load", load)

	result := true

	// get all the loads for the same week as the requested load
	loads := lh.db.GetLoads(load.CustomerID, true, utils.GetStartOfWeek(load.Time), utils.GetEndOfWeek(load.Time))
	lh.l.Infof("found %d loads", len(loads))
	lh.l.Infof("start time %s end time %s", utils.GetStartOfWeek(load.Time), utils.GetEndOfWeek(load.Time))

	weeklySum, err := lh.addLoadAmounts(loads)

	if err != nil {
		// defensive coding
		// input validation should prevent this
		// if we could not calculate the weekly sum, do not accept the load
		lh.l.Errorln("unable to calculate weekly sum", err)
		return false
	}

	amount, err := utils.ConvertLoadAmount(load.LoadAmount)
	if err != nil {
		// defensive coding
		// input validation should prevent this 
		// if we could not get the amount, do not accept the load
		lh.l.Errorln("unable to get amount from fundsload request", err)
		return false
	}

	if (weeklySum + amount) > maxLoadAmountPerWeek {
		result = false
		lh.l.Infoln("exceeded maximum load amount per week")
	}

	return result
}

// AddLoadAmounts calculates the sum for a set of loads
func (lh *Loads) addLoadAmounts(loads []*data.Load) (uint32, error) {
	sum := uint32(0)

	for _, load := range loads {

		lh.l.Infoln("adding for load", load)
		amount, err := utils.ConvertLoadAmount(load.LoadAmount)
		if err != nil {
			// defensive coding
			// input validation should prevent this 
			lh.l.Errorln("unable to get amount from fundsload", err)
			return 0, err
		}
		sum = sum + amount
	}
	return sum, nil
}
