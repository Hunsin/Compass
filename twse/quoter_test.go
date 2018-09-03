package twse

import (
	"testing"
	"time"

	"github.com/Hunsin/compass/trade"
)

var (
	qt = &quoter{}

	tsmc20171016 = trade.Quote{
		Date:   "2017-10-16",
		Open:   237.5,
		High:   238,
		Low:    236,
		Close:  238,
		Volume: 20673072,
		Avg:    237.360144395,
	}

	largan20171106 = trade.Quote{
		Date:   "2017-11-06",
		Open:   5950,
		High:   5950,
		Low:    5825,
		Close:  5875,
		Volume: 463126,
		Avg:    5889.290171573179,
	}
)

func TestMonth(t *testing.T) {

	// not listed
	q, err := qt.Month("2330", 1994, time.August)
	if err == nil {
		t.Error("quoter.Month failed: did not return error when the security is unlisted")
	}

	q, err = qt.Month("2330", 2017, time.October)
	if err != nil {
		t.Fatal("quoter.Month failed:", err)
	}

	// 19 days opened in Oct. 2017
	if len(q) != 19 {
		t.Errorf("quoter.Month failed: didn't return complete data\nNum of quotes: %d", len(q))
	}

	for i := range q {
		if equal(q[i], trade.Quote{}) {
			t.Error("quoter.Month failed: an empty quote was returned")
		}
	}
	if !equal(q[7], tsmc20171016) {
		t.Errorf("quoter.Month failed: worng data\nGot : %v\nWant: %v", q[7], tsmc20171016)
	}
}

func TestYear(t *testing.T) {
	days := map[int]int{
		1994: 93,  //  93 days opened in 1994
		2016: 244, // 244 days opened in 2016
	}

	for y := range days {
		q, err := qt.Year("2330", y)
		if err != nil {
			t.Fatal("quoter.Year failed:", err)
		}

		if len(q) != days[y] {
			t.Errorf("quoter.Year failed: didn't return complete data\nNum. of quotes: %d, Want: %d", len(q), days[y])
		}
	}
}
