package crawler

import (
	"testing"
	"time"
)

var (
	d1 = Date{2010, time.March, 10}
	d2 = Date{2017, time.November, 30}
)

func TestAfter(t *testing.T) {
	if d1.After(d2) {
		t.Error("d1 should not after d2")
	}
	if d1.After(d1) {
		t.Error("d1 should not after d1")
	}
}

func TestBefore(t *testing.T) {
	if d2.Before(d1) {
		t.Error("d2 should not before d1")
	}
	if d2.After(d2) {
		t.Error("d2 should not after d2")
	}
}
