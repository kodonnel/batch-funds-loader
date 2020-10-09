package utils

import "time"

// GetStartOfDay returns a time object for the start of the day
// time can be provided in any timezone
// Each day is considered to end at midnight UTC
// result is returned in the timezone of the request param
func GetStartOfDay(t time.Time) time.Time {

	loc := t.Location()
	utc, _ := time.LoadLocation("UTC")

	// convert to UTC
	t = t.In(utc)

	// get start of the day
	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, utc)

	// convert back to original timezone
	start = start.In(loc)
	return start
}

// GetEndOfDay returns a time object for the end of the day
// time can be provided in any timezone
// Each day is considered to end at midnight UTC
// result is returned in the timezone of the request param
func GetEndOfDay(t time.Time) time.Time {
	loc := t.Location()
	utc, _ := time.LoadLocation("UTC")

	// convert to UTC
	t = t.In(utc)

	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, utc)

	// convert back to original timezone
	end = end.In(loc)
	return end
}

// GetStartOfWeek returns a time object for the start of the day
// time can be provided in any timezone
// week starts on Monday (i.e. one second after 23:59:59 on Sunday) UTC
// result is returned in the timezone of the request param
func GetStartOfWeek(t time.Time) time.Time {
	loc := t.Location()
	utc, _ := time.LoadLocation("UTC")

	// convert to UTC
	t = t.In(utc)

	daysToRemove := 0
	switch weekday := int(t.Weekday()); weekday {
	case 0:
		daysToRemove = 6 // sunday
	case 1:
		daysToRemove = 0 // monday
	case 2:
		daysToRemove = 1 // tuesday
	case 3:
		daysToRemove = 2 // wednesday
	case 4:
		daysToRemove = 3 // thursday
	case 5:
		daysToRemove = 4 // friday
	case 6:
		daysToRemove = 5 // saturday
	}
	newT := t.AddDate(0, 0, (-1 * daysToRemove))
	start := time.Date(newT.Year(), newT.Month(), newT.Day(), 0, 0, 0, 0, utc)

	// convert back to original timezone
	start = start.In(loc)
	return start
}

// GetEndOfWeek returns a time object for the end of the day
// time can be provided in any timezone
// Week ends  23:59:59 on Sunday UTC
// result is returned in the timezone of the request param
func GetEndOfWeek(t time.Time) time.Time {
	loc := t.Location()
	utc, _ := time.LoadLocation("UTC")

	// convert to UTC
	t = t.In(utc)

	daysToAdd := 0
	switch weekday := int(t.Weekday()); weekday {
	case 0:
		daysToAdd = 0 // sunday
	case 1:
		daysToAdd = 6 // monday
	case 2:
		daysToAdd = 5 // tuesday
	case 3:
		daysToAdd = 4 // wednesday
	case 4:
		daysToAdd = 3 // thursday
	case 5:
		daysToAdd = 2 // friday
	case 6:
		daysToAdd = 1 // saturday
	}
	newT := t.AddDate(0, 0, daysToAdd)
	end := time.Date(newT.Year(), newT.Month(), newT.Day(), 23, 59, 59, 999999999, utc)

	// convert back to original timezone
	end = end.In(loc)
	return end
}

// IsInTimeSpan returns true if the given time is between the start and end times
func IsInTimeSpan(start, end, check time.Time) bool {

	// get each time in utc
	utc, _ := time.LoadLocation("UTC")

	sutc := start.In(utc)
	eutc := end.In(utc)
	cutc := check.In(utc)

	// check inclusive between
	return cutc.After(sutc.Add(time.Duration(-1)*time.Second)) && cutc.Before(eutc.Add(time.Duration(1)*time.Second))
}
