package meetup

import (
	"time"
)

type WeekSchedule int

const (
	First WeekSchedule = 1 // Day in the week
	Second = 8 // One week on
	Third = 15 // Two weeks on
	Fourth = 22 // Three weeks on
	Teenth = 13 // Middle of the month (13 to support February)
	Last = -1
)

// Day returns the date for a given meetup date
// This problem was badly worded and it took me a while
// inspired from https://exercism.io/tracks/javascript/exercises/meetup/solutions/b23b7fdeb7fe4b59b4cc3c4861845fe7
func Day(week WeekSchedule, day time.Weekday, month time.Month, year int) int {
	var startDate time.Time
	switch week {
	case First, Second, Teenth, Third, Fourth:
		startDate = time.Date(year, month, int(week), 0, 0, 0, 0, time.UTC)
	default:
		// Get the date for the next month
		startDate = time.Date(year, month + 1, 1, 0, 0, 0, 0, time.UTC)
		dayOfMonth := 1 + (int(day) - int(startDate.Weekday()) + 7)%7
		nextMonth := time.Date(year, month + 1, dayOfMonth, 0, 0, 0, 0, time.UTC)

		// Now subtract 7 days
		what := nextMonth.AddDate(0, 0, -7)
		return what.Day()
	}

	// Think about this, it is from the start date day plus one week but we can only have weeks (mod 7)
	// so then we add the day of the week we want (because we have the right amount of days)
	return int(week) + (int(day) - int(startDate.Weekday()) + 7)%7
}
