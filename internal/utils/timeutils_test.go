package utils

import (
	"testing"
	"time"
)

// scenario
// input is the start of the day
func TestGetStartOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	testTime := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)
	expected := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	result := GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day
func TestGetStartOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	testTime := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	// expected = 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	expected := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	result := GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the start of the day
func TestGetEndOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	testTime := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	expected := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	result := GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day
func TestGetEndOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 19:59 oct 7 Toronto time = 23:59 oct 8 UTC
	testTime := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	expected := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	result := GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a monday (beginning of week)
func TestGetStartOfWeekWithMonday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	testTime := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a saturday
func TestGetStartOfWeekWithSaturday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 19:59 oct 03 Toronto time = 23:59 oct 03 UTC (Sat)
	testTime := time.Date(2020, 10, 3, 19, 59, 59, 999999999, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Sat)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a sunday (very end of week)
func TestGetStartOfWeekWithSunday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 19:59 oct 04 Toronto time = 23:59 oct 04 UTC (Sun)
	testTime := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Sat)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a sunday
func TestGetEndOfWeekWithSunday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 26 Toronto time = 00:00 sept 27 UTC (Sun)
	testTime := time.Date(2020, 9, 26, 20, 0, 0, 0, loc)

	// 19:59 sept 27 Toronto time = 23:59 sept 27 UTC (Sun)
	expected := time.Date(2020, 9, 27, 19, 59, 59, 999999999, loc)
	result := GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a saturday
func TestGetEndOfWeekWithSaturday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 19:59 oct 03 Toronto time = 23:59 oct 03 UTC (Sat)
	testTime := time.Date(2020, 10, 3, 19, 59, 59, 999999999, loc)

	// 19:59 oct 04 Toronto time = 23:59 oct 04 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)

	result := GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// time is within the start and end
func TestIsInTimeSpanWithTimeInRange(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	start := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	end := time.Date(2020, 10, 4, 21, 59, 59, 999999999, loc)
	check := time.Date(2020, 10, 4, 21, 59, 59, 999999998, loc)
	expected := true

	result := IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// time is after the start and end
func TestIsInTimeSpanWithTimeAfterRange(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	start := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	end := time.Date(2020, 10, 4, 21, 59, 59, 999999999, loc)
	check := time.Date(2020, 10, 4, 22, 0, 0, 0, loc)
	expected := true

	result := IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// time is before the start and end
func TestIsInTimeSpanWithTimeBeforeRange(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	start := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	end := time.Date(2020, 10, 4, 21, 59, 59, 999999999, loc)
	check := time.Date(2020, 10, 3, 23, 59, 59, 999999998, loc)
	expected := true

	result := IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}
