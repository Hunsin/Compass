package crawler

import (
	"testing"
	"time"
)

var (
	d1 = Date{2001, time.March, 5}
	d2 = Date{2009, time.November, 15}
	f1 = []string{
		"2001/03/05",
		"2001-03-05",
		"Mar 05 2001",
	}
	f2 = []string{
		"2009/11/15",
		"2009-11-15",
		"15 Nov 2009",
	}
)

func TestParseDate(t *testing.T) {
	layouts := []string{
		"2006/01/02",
		"2006-01-02",
		"Jan 2 2006",
	}
	for i, l := range layouts {
		d, err := ParseDate(l, f1[i])
		if err != nil {
			t.Errorf("ParseDate exit with error: %v", err)
		}
		if d.Year != d1.Year ||
			d.Month != d1.Month ||
			d.Day != d1.Day {
			t.Errorf("ParseDate failed: parse %s returns %v", f1[i], d)
		}
	}

	// Invalid date should return error
	_, err := ParseDate(layouts[0], "2017/13/01")
	if err == nil {
		t.Error("ParseDate failed: Invalid input doesn't return error")
	}
}

func TestAfter(t *testing.T) {
	if d1.After(d2) {
		t.Error("Date.After failed: d1 should not after d2")
	}
	if d1.After(d1) {
		t.Error("Date.After failed: d1 should not after d1")
	}
}

func TestBefore(t *testing.T) {
	if d2.Before(d1) {
		t.Error("Date.Before failed: d2 should not before d1")
	}
	if d2.After(d2) {
		t.Error("Date.Before failed: d2 should not after d2")
	}
}

func TestEqual(t *testing.T) {
	if d1.Equal(d2) {
		t.Error("Date.Equal failed: d1 != d2")
	}
	if !d2.Equal(d2) {
		t.Error("Date.Equal failed: d2 == d2")
	}
}

func TestMarshalJSON(t *testing.T) {
	b, err := d1.MarshalJSON()
	if err != nil {
		t.Errorf("Date.MarshalJSON exits with error: %v", err)
	}

	w := `"` + f1[0] + `"`
	if string(b) != w {
		t.Errorf("Date.MarshalJSON failed. want: %s, got: %s", w, string(b))
	}
}

func TestMarshalText(t *testing.T) {
	b, err := d1.MarshalText()
	if err != nil {
		t.Errorf("Date.MarshalText exits with error: %v", err)
	}

	w := f1[0]
	if string(b) != w {
		t.Errorf("Date.MarshalText failed. want: %s, got: %s", w, string(b))
	}
}

func TestString(t *testing.T) {
	w := f2[0]
	if d2.String() != w {
		t.Errorf("Date.String failed. want: %s, got: %v", w, d2)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	var d Date
	for _, b := range f2 {
		err := d.UnmarshalJSON([]byte(`"` + b + `"`))
		if err != nil {
			t.Errorf("Date.UnmarshalJSON exits with error: %v", err)
		}
		if !d.Equal(d2) {
			t.Errorf("Date.UnmarshalJSON failed. want: %v, got: %v", d2, d)
		}
	}
}

func TestUnmarshalText(t *testing.T) {
	var d Date
	for _, b := range f2 {
		err := d.UnmarshalText([]byte(b))
		if err != nil {
			t.Errorf("Date.UnmarshalText exits with error: %v", err)
		}
		if !d.Equal(d2) {
			t.Errorf("Date.UnmarshalText failed. want: %v, got: %v", d2, d)
		}
	}
}
