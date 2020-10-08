package handlers

import (
	"github.com/kodonnel/batch-funds-loader/internal/data"
	"github.com/kodonnel/batch-funds-loader/internal/utils"
	"github.com/sirupsen/logrus"
)

// Loads handler for getting and updating funds load requests
type Loads struct {
	l  *logrus.Logger
	db *data.LoadsDB
}

// NewLoads returns a new products handler with the given logger
func NewLoads(l *logrus.Logger, db *data.LoadsDB) *Loads {
	return &Loads{l, db}
}

// ProcessLoadRequest processes the load request and sends a response through a receiving channel
func (lh *Loads) ProcessLoadRequest(req data.Load, msg chan<- data.LoadResult) {

	lh.l.Infoln("processing load request", req)

	// check if duplicate
	if !lh.db.IsDuplicate(req) {

		sameDayLoads := lh.db.GetLoads(req.CustomerID, true, utils.GetStartForDay(req.Time), utils.GetEndForDay(req.Time))
		accepted := true

		if len(sameDayLoads) >= 3 {
			accepted = false
		}

		dailySum := 0.0
		for _, existingLoad := range sameDayLoads {
			dailySum = dailySum + utils.GetFloatAmount(existingLoad.LoadAmount)
		}
		if (dailySum + utils.GetFloatAmount(req.LoadAmount)) >= 5000 {
			accepted = false
		}

		sameWeekLoads := lh.db.GetLoads(req.CustomerID, true, utils.GetStartForWeek(req.Time), utils.GetEndForWeek(req.Time))

		weeklySum := 0.0
		for _, existingWeekLoad := range sameWeekLoads {
			weeklySum = weeklySum + utils.GetFloatAmount(existingWeekLoad.LoadAmount)
		}

		if (weeklySum + utils.GetFloatAmount(req.LoadAmount)) >= 20000 {
			accepted = false
		}

		req.Accepted = accepted // replace with IsValidLoad function
		lh.db.AddLoad(req)

		var loadResult *data.LoadResult
		loadResult = new(data.LoadResult)
		loadResult.CustomerID = req.CustomerID
		loadResult.ID = req.ID
		loadResult.Accepted = req.Accepted

		msg <- *loadResult
	}

}
