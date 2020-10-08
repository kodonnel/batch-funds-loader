package handlers

import (
	"fmt"

	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/kodonnel/batch-funds-loader/internal/utils"
	"github.com/sirupsen/logrus"
)

// ErrDuplicateFound indicates that a load with the same id and customer_id was already processed
var ErrDuplicateFound = fmt.Errorf("Duplicate Load Request Found")

// Loads handler for getting and updating funds load requests
type Loads struct {
	l  *logrus.Logger
	db *data.LoadsDB
}

// NewLoads returns a new products handler with the given logger
func NewLoads(l *logrus.Logger, db *data.LoadsDB) *Loads {
	return &Loads{l, db}
}

// ProcessLoadRequest processes the load request and return the result
// (apply business logic)
func (lh *Loads) ProcessLoadRequest(req data.Load) (*data.LoadResult, error) {

	lh.l.Infoln("processing load request", req)

	var loadResult *data.LoadResult

	// check if duplicate
	if !lh.db.IsDuplicate(req) {

		req.Accepted = (lh.isWithinDailyLimits(req) && lh.isWithinWeeklyLimits(req))

		lh.db.AddLoad(req)

		loadResult = new(data.LoadResult)
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

	sameDayLoads := lh.db.GetLoads(load.CustomerID, true, utils.GetStartOfDay(load.Time), utils.GetEndOfDay(load.Time))

	if len(sameDayLoads) >= 3 {
		lh.l.Infoln("exceeded maximum loads per day")
		result = false
	}

	dailySum := 0.0
	for _, existingLoad := range sameDayLoads {

		eamount, err := utils.GetFloatAmount(existingLoad.LoadAmount)
		if err != nil {
			lh.l.Errorln("unable to get amount from fundsload", err)
			return false
		}

		dailySum = dailySum + eamount
	}

	ramount, err := utils.GetFloatAmount(load.LoadAmount)

	if err != nil {
		lh.l.Errorln("unable to get amount from fundsload request", err)
		return false
	}

	if (dailySum + ramount) >= 5000 {
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

	sameWeekLoads := lh.db.GetLoads(load.CustomerID, true, utils.GetStartOfWeek(load.Time), utils.GetEndOfWeek(load.Time))

	weeklySum := 0.0
	for _, existingWeekLoad := range sameWeekLoads {

		weamount, err := utils.GetFloatAmount(existingWeekLoad.LoadAmount)

		if err != nil {
			lh.l.Errorln("unable to get amount from fundsload request", err)
			return false
		}
		weeklySum = weeklySum + weamount
	}

	rwamount, err := utils.GetFloatAmount(load.LoadAmount)
	if err != nil {
		// if we could not get the amount, do not accept the transaction
		lh.l.Errorln("unable to get amount from fundsload request", err)
		return false
	}

	if (weeklySum + rwamount) >= 20000 {
		result = false
		lh.l.Infoln("exceeded maximum load amount per week")

	}

	return result
}
