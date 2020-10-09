package utils_test

import (
	"testing"
	"time"

	"github.com/kodonnel/batch-funds-loader/internal/utils"
)

// scenario
// input is the start of the day (UTC)
func TestGetStartOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	testTime := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)
	expected := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	result := utils.GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day (UTC)
func TestGetStartOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	testTime := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	// expected = 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	expected := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	result := utils.GetStartOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the start of the day (UTC)
func TestGetEndOfDayWithStartOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 20:00 oct 7 Toronto time = 00:00 oct 8 UTC
	testTime := time.Date(2020, 10, 7, 20, 0, 0, 0, loc)

	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	expected := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	result := utils.GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is the end of the day (UTC)
func TestGetEndOfDayWithEndOfDay(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// use America/Toronto timezone (UTC-4)
	// 19:59 oct 7 Toronto time = 23:59 oct 8 UTC
	testTime := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	// 19:59 oct 8 Toronto time = 23:59 oct 8 UTC
	expected := time.Date(2020, 10, 8, 19, 59, 59, 999999999, loc)

	result := utils.GetEndOfDay(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a monday (beginning of week) (UTC)
func TestGetStartOfWeekWithMonday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	testTime := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := utils.GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a tuesday
func TestGetStartOfWeekWithTuesday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 28 Toronto time = 00:00 sept 29 UTC (Tue)
	testTime := time.Date(2020, 9, 28, 20, 12, 34, 123, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := utils.GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a wednesday
func TestGetStartOfWeekWithWednesday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 29 Toronto time = 00:00 sept 30 UTC (Wed)
	testTime := time.Date(2020, 9, 29, 20, 12, 34, 123, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := utils.GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a thursday
func TestGetStartOfWeekWithThursday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 30 Toronto time = 00:00 oct 01 UTC (Thu)
	testTime := time.Date(2020, 9, 30, 20, 12, 34, 123, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := utils.GetStartOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a friday
func TestGetStartOfWeekWithFriday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 oct 01 Toronto time = 00:00 oct 02 UTC (Fri)
	testTime := time.Date(2020, 10, 1, 20, 12, 34, 123, loc)

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	expected := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	result := utils.GetStartOfWeek(testTime)
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

	result := utils.GetStartOfWeek(testTime)
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

	result := utils.GetStartOfWeek(testTime)
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
	result := utils.GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a monday
func TestGetEndOfWeekWithMonday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 27 Toronto time = 00:00 sept 28 UTC (Mon)
	testTime := time.Date(2020, 9, 27, 20, 0, 0, 0, loc)

	// 19:59 oct 4 27 Toronto time = 23:59 oct 4 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)
	result := utils.GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a tuesday
func TestGetEndOfWeekWithTuesday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 28 Toronto time = 00:00 sept 29 UTC (Tue)
	testTime := time.Date(2020, 9, 28, 20, 0, 0, 0, loc)

	// 19:59 oct 4 27 Toronto time = 23:59 oct 4 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)
	result := utils.GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a wednesday
func TestGetEndOfWeekWithWednesday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 29 Toronto time = 00:00 sept 30 UTC (Wed)
	testTime := time.Date(2020, 9, 29, 20, 0, 0, 0, loc)

	// 19:59 oct 4 27 Toronto time = 23:59 oct 4 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)
	result := utils.GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a thursday
func TestGetEndOfWeekWithThursday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 sept 30 Toronto time = 00:00 oct 1 UTC (Thurs)
	testTime := time.Date(2020, 9, 30, 20, 0, 0, 0, loc)

	// 19:59 oct 4 27 Toronto time = 23:59 oct 4 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)
	result := utils.GetEndOfWeek(testTime)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// input is a friday
func TestGetEndOfWeekWithFriday(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	// 20:00 oct 1 Toronto time = 00:00 oct 2 UTC (Fri)
	testTime := time.Date(2020, 10, 1, 20, 0, 0, 0, loc)

	// 19:59 oct 4 27 Toronto time = 23:59 oct 4 UTC (Sun)
	expected := time.Date(2020, 10, 4, 19, 59, 59, 999999999, loc)
	result := utils.GetEndOfWeek(testTime)
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

	result := utils.GetEndOfWeek(testTime)
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

	result := utils.IsInTimeSpan(start, end, check)
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

	result := utils.IsInTimeSpan(start, end, check)
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

	result := utils.IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// time, start and end are the same
func TestIsInTimeSpanWithStartAndEnd(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	start := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	end := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	check := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	expected := true

	result := utils.IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}

// scenario
// time is before the start and end
// start and end are the same
func TestIsInTimeSpanWithStartAndEndBefore(t *testing.T) {
	loc, _ := time.LoadLocation("America/Toronto")

	start := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	end := time.Date(2020, 10, 3, 23, 59, 59, 999999999, loc)
	check := time.Date(2020, 10, 3, 23, 59, 58, 0, loc)
	expected := false

	result := utils.IsInTimeSpan(start, end, check)
	if result != expected {
		t.Errorf("failed expected %v got %v", expected, result)
	}
}