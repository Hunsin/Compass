package twse

import (
	"crawler"
	"math"
	"testing"
	"time"
)

var (
	tw   crawler.Market
	tsmc crawler.Security
)

func init() {
	tw, _ = crawler.Open("twse")
	tsmc, _ = tw.Search("2330")
}

func compare(want, got crawler.Daily) bool {
	return (want.Date.Equal(got.Date) &&
		want.Open == got.Open &&
		want.High == got.High &&
		want.Low == got.Low &&
		want.Close == got.Close &&
		want.Volume == got.Volume &&
		math.Abs(want.Avg-got.Avg) > 0.00001)
}

func TestOpen(t *testing.T) {
	var err error
	tw, err = crawler.Open("twse")
	if err != nil {
		t.Fatalf("crawler.Open exits with error: %v", err)
	}
}

func TestSearchAndSymbol(t *testing.T) {
	samples := map[string]bool{
		"0000": false,
		"2330": true,
	}

	for w := range samples {
		st, err := tw.Search(w)
		if err != nil {
			if !samples[w] {
				continue
			}
			t.Fatalf("Search %s exits with error: %v", w, err)
		}

		if !samples[w] {
			t.Errorf("Search %s doesn't return not found error", w)
		}

		if got := st.Symbol(); got != w {
			t.Errorf("Symbol() not match, want: %s, got: %s", w, got)
		}
	}
}

func TestName(t *testing.T) {
	samples := map[string]string{
		"2330":   "TSMC",    // Stock
		"03003X": "YX 03",   // Listing Warrants
		"9103":   "MEDTECS", // TDR
	}

	for k := range samples {
		st, err := tw.Search(k)
		if err != nil {
			t.Fatalf("Search %s exits with error: %v", k, err)
		}

		if got := st.Name(); got != samples[k] {
			t.Errorf("Name() not match, want: %s, got: %s", samples[k], got)
		}
	}
}

func TestDate(t *testing.T) {
	dt, _ := time.ParseInLocation(dateFormat, "2017/11/06 14:30", cst)

	// 2017/11/06 data of TSMC(2330) and LARGAN(3008)
	samples := map[string]crawler.Daily{
		"2330": crawler.Daily{
			Date:   dt,
			Open:   243.5,
			High:   244,
			Low:    239,
			Close:  239.5,
			Volume: 21029515,
			Avg:    240.78876284117823,
		},
		"3008": crawler.Daily{
			Date:   dt,
			Open:   5950,
			High:   5950,
			Low:    5825,
			Close:  5875,
			Volume: 463126,
			Avg:    5889.290171573179,
		},
	}

	for k := range samples {
		st, err := tw.Search(k)
		if err != nil {
			t.Fatalf("Search %s exits with error: %v", k, err)
		}

		// This date is closed
		got, err := st.Date(2017, 11, 4)
		if err == nil {
			t.Error("Stock.Date doesn't return error in closed date 2017/11/4")
		}

		got, err = st.Date(2017, 11, 6)
		if err != nil {
			t.Errorf("Stock.Date exists with error: %v", err)
		}

		if !compare(samples[k], got) {
			t.Errorf("Stock.Date failed\nGot : %v\nWant: %v", got, samples[k])
		}
	}
}

func TestMonth(t *testing.T) {
	got, err := tsmc.Month(2017, 10)
	if err != nil {
		t.Errorf("Stock.Month exists with error: %v", err)
	}

	// There were 19 days opened in Oct. 2017
	if len(got) != 19 {
		t.Errorf("Stock.Month doesn't return complete data\nNum. of Daily: %d", len(got))
	}
}

func TestYear(t *testing.T) {
	got, err := tsmc.Year(2016)
	if err != nil {
		t.Errorf("Stock.Year exists with error: %v", err)
	}

	// There were 244 days opened in 2016
	if len(got) != 244 {
		t.Errorf("Stock.Year doesn't return complete data\nNum. of Daily: %d", len(got))
	}
}
