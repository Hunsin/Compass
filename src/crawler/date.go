package crawler

import (
	"fmt"
	"time"
)

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

// Equal reports whether d and t are the same.
func (d Date) Equal(t Date) bool {
	return !d.After(t) && !d.Before(t)
}

// MarshalJSON implements the json.Marshaler interface.
// The output is in "2006/01/02" format.
func (d Date) MarshalJSON() ([]byte, error) {
	s, _ := d.MarshalText()
	b := make([]byte, 0, len(s)+2)
	b = append(b, '"')
	b = append(b, s...)
	return append(b, '"'), nil
}

// MarshalText implements the encoding.TextMarshaler interface.
// The output is in "2006/01/02" format.
func (d Date) MarshalText() ([]byte, error) {
	s := []byte(d.String())
	b := make([]byte, 0, len(s))
	return append(b, s...), nil
}

// Scan implements the database/sql.Scanner interface.
func (d *Date) Scan(v interface{}) error {
	switch s := v.(type) {
	case time.Time:
		*d = Date{Year: s.Year(), Month: s.Month(), Day: s.Day()}
		return nil
	case []byte:
		return d.UnmarshalText(s)
	case string:
		return d.UnmarshalText([]byte(s))
	}
	return fmt.Errorf("crawler.Date: Unsupport scanning type %T", v)
}

// String returns a string of date in "2006/01/02" format.
func (d Date) String() string {
	return fmt.Sprintf("%4d/%02d/%02d", d.Year, d.Month, d.Day)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
// The formats it supports are "2006/01/02", "2006-01-02" and "02 Jan 2006".
func (d *Date) UnmarshalJSON(b []byte) error {
	if string(b) == "null" {
		return nil
	}

	return d.UnmarshalText(b[1 : len(b)-1])
}

// UnmarshalText implements the encoding.TextUnmarshaler interface.
// The formats it supports are "2006/01/02", "2006-01-02" and "02 Jan 2006".
func (d *Date) UnmarshalText(b []byte) error {
	var t time.Time
	var err error

	s := string(b)
	for _, l := range []string{"2006/01/02", "2006-01-02", "02 Jan 2006"} {
		t, err = time.Parse(l, s)
		if err == nil {
			break
		}
	}
	if err != nil {
		return err
	}

	*d = Date{t.Year(), t.Month(), t.Day()}
	return nil
}
