package crawler

import (
	"testing"
	"time"
)

var (
	d1 = Date{2001, time.March, 5}
	d2 = Date{2009, time.November, 15}
)

func TestParseDate(t *testing.T) {
	layouts := map[string]string{
		"2006/01/02": "2001/03/05",
		"2006-1-02":  "2001-3-05",
		"Jan 2 2006": "Mar 05 2001",
	}
	for l, s := range layouts {
		d, err := ParseDate(l, s)
		if err != nil {
			t.Errorf("ParseDate exit with error: %v", err)
		}
		if d.Year != d1.Year ||
			d.Month != d1.Month ||
			d.Day != d1.Day {
			t.Errorf("ParseDate failed: parse %s returns %v", d)
		}
	}

	// Invalid date should return error
	_, err := ParseDate("2006/01/02", "2017/13/01")
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

	w := `"2001/03/05"`
	if string(b) != w {
		t.Errorf("Date.MarshalJSON failed. want: %s, got: %s", w, string(b))
	}
}

func TestString(t *testing.T) {
	w := "2009/11/15"
	if d2.String() != w {
		t.Errorf("Date.String failed. want: %s, got: %v", w, d2)
	}
}

func TestUnmarshalJSON(t *testing.T) {
	bs := [][]byte{
		[]byte(`"2009-11-15"`),
		[]byte(`"2009/11/15"`),
		[]byte(`"15 Nov 2009"`),
	}

	var d Date
	for _, b := range bs {
		err := d.UnmarshalJSON(b)
		if err != nil {
			t.Errorf("Date.UnmarshalJSON exits with error: %v", err)
		}
		if !d.Equal(d2) {
			t.Errorf("Date.UnmarshalJSON failed. want: %v, got: %v", d2, d)
		}
	}
}
