package utils

import "time"

// GetStartOfDay returns a time object for the start of the day
func GetStartOfDay(t time.Time) time.Time {

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return start
}

// GetEndOfDay returns a time object for the end of the day
func GetEndOfDay(t time.Time) time.Time {

	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return end
}

// GetStartOfWeek returns a time object for the start of the day
func GetStartOfWeek(t time.Time) time.Time {

	weekday := int(t.Weekday())
	daysToRemove := weekday * -1
	newT := t.AddDate(0, 0, daysToRemove)
	start := time.Date(newT.Year(), newT.Month(), newT.Day(), 0, 0, 0, 0, newT.Location())
	return start
}

// GetEndOfWeek returns a time object for the end of the day
func GetEndOfWeek(t time.Time) time.Time {

	weekday := int(t.Weekday())
	daysToAdd := 6 - weekday
	newT := t.AddDate(0, 0, daysToAdd)
	end := time.Date(newT.Year(), newT.Month(), newT.Day(), 23, 59, 59, 999999999, newT.Location())
	return end
}

// IsInTimeSpan returns true if the given time is between the start and end times
func IsInTimeSpan(start, end, check time.Time) bool {

	// check inclusive between
	return check.After(start.Add(time.Duration(-1)*time.Second)) && check.Before(end.Add(time.Duration(1)*time.Second))
}
