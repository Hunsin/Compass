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
		math.Abs(want.Avg-got.Avg) < 0.00001)
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

func TestProperties(t *testing.T) {
	samples := map[string][]string{
		"2330":   []string{"TSMC", "Stocks", "1994/09/05 09:00"},            // Stock
		"0052":   []string{"FB TECHNOLOGY", "ETF", "2006/09/12 09:00"},      // ETF
		"03003X": []string{"YX 03", "Listing Warrants", "2014/07/31 09:00"}, // Listing Warrants
		"9103":   []string{"MEDTECS", "TDR", "2002/12/13 09:00"},            // TDR
	}

	for k := range samples {
		st, err := tw.Search(k)
		if err != nil {
			t.Fatalf("Search %s exits with error: %v", k, err)
		}

		if got := st.Market(); got != "TWSE" {
			t.Errorf("Market() not match, want: TWSE, got: %s", got)
		}

		if got := st.Name(); got != samples[k][0] {
			t.Errorf("Name() not match, want: %s, got: %s", samples[k][0], got)
		}

		if got := st.Type(); got != samples[k][1] {
			t.Errorf("Type() not match, want: %s, got: %s", samples[k][1], got)
		}

		d, _ := time.ParseInLocation(dateFormat, samples[k][2], cst)
		if got := st.Listed(); !got.Equal(d) {
			t.Errorf("Listed() not match, want: %v, got: %v", samples[k][2], got)
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

	// Not listed
	got, err := tsmc.Month(1994, 8)
	if _, ok := err.(*crawler.ErrNotListed); !ok {
		t.Errorf("Stock.Month doesn't return error with the month not listed")
	}

	got, err = tsmc.Month(2017, 10)
	if err != nil {
		t.Errorf("Stock.Month exists with error: %v", err)
	}

	// There were 19 days opened in Oct. 2017
	if len(got) != 19 {
		t.Errorf("Stock.Month doesn't return complete data\nNum. of Daily: %d", len(got))
	}
}

func TestYear(t *testing.T) {
	samples := map[int]int{
		1994: 93,  // There were  93 days opened in 1994
		2016: 244, // There were 244 days opened in 2016
	}

	for y := range samples {
		got, err := tsmc.Year(y)
		if err != nil {
			t.Errorf("Stock.Year exists with error: %v", err)
		}

		if len(got) != samples[y] {
			t.Errorf("Stock.Year doesn't return complete data\nNum. of Daily: %d, Want: %d", len(got), samples[y])
		}
	}
}
