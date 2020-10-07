package utils

import "time"

// GetStartForDay returns a time object for the start of the day
func GetStartForDay(t time.Time) time.Time {

	start := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
	return start
}

// GetEndForDay returns a time object for the end of the day
func GetEndForDay(t time.Time) time.Time {

	end := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 999999999, t.Location())
	return end
}

// GetStartForWeek returns a time object for the start of the day
func GetStartForWeek(t time.Time) time.Time {

	weekday := int(t.Weekday())
	daysToRemove := weekday * -1
	newT := t.AddDate(0, 0, daysToRemove)
	start := time.Date(newT.Year(), newT.Month(), newT.Day(), 0, 0, 0, 0, newT.Location())
	return start
}

// GetEndForWeek returns a time object for the end of the day
func GetEndForWeek(t time.Time) time.Time {

	weekday := int(t.Weekday())
	daysToAdd := 6 - weekday
	newT := t.AddDate(0, 0, daysToAdd)
	end := time.Date(newT.Year(), newT.Month(), newT.Day(), 23, 59, 59, 999999999, newT.Location())
	return end
}

// InTimeSpan returns true if the given time is between the start and end times
func InTimeSpan(start, end, check time.Time) bool {

	// check inclusive between
	return check.After(start.Add(time.Duration(-1)*time.Second)) && check.Before(end.Add(time.Duration(1)*time.Second))
}
