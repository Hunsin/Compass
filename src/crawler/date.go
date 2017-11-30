package crawler

import "time"

// A Date specifies the year, month and day.
type Date struct {
	Year  int
	Month time.Month
	Day   int
}

// ParseDate parses the d with layout and returns the value of Date.
// The layout follows the format of time.Parse.
func ParseDate(layout, d string) (Date, error) {
	t, err := time.Parse(layout, d)
	if err != nil {
		return Date{}, err
	}

	return Date{t.Year(), t.Month(), t.Day()}, nil
}

// After reports whether d is after t.
func (d Date) After(t Date) bool {
	if d.Year != t.Year {
		return d.Year > t.Year
	}
	if d.Month != t.Month {
		return d.Month > t.Month
	}
	return d.Day > t.Day
}

// Before reports whether d is before t.
func (d Date) Before(t Date) bool {
	return t.After(d)
}
