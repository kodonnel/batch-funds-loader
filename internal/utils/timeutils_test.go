package utils

import (
	"testing"
	"time"
)

// scenario
// input is the start of the day
func TestGetStartOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 8, 0, 0, 0, 0, loc)
	expected := time.Date(2020, 10, 8, 0, 0, 0, 0, loc)

	result := GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day
func TestGetStartOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 8, 23, 59, 59, 999999999, loc)
	expected := time.Date(2020, 10, 8, 0, 0, 0, 0, loc)

	result := GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the start of the day
func TestGetEndOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 8, 0, 0, 0, 0, loc)
	expected := time.Date(2020, 10, 8, 23, 59, 59, 999999999, loc)

	result := GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day
func TestGetEndOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 8, 23, 59, 59, 999999999, loc)
	expected := time.Date(2020, 10, 8, 23, 59, 59, 999999999, loc)

	result := GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a sunday
func TestGetStartOfWeekWithSunday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 9, 27, 0, 0, 0, 0, loc)
	expected := time.Date(2020, 9, 27, 0, 0, 0, 0, loc)

	result := GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a saturday
func TestGetStartOfWeekWithSaturday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	expected := time.Date(2020, 9, 27, 0, 0, 0, 0, loc)

	result := GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a sunday
func TestGetEndOfWeekWithSunday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 9, 27, 0, 0, 0, 0, loc)
	expected := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	result := GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a saturday
func TestGetEndOfWeekWithSaturday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	testTime := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	expected := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)

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
