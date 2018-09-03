package twse

import (
	"math"
	"reflect"
	"testing"

	"github.com/Hunsin/compass/market"
	"github.com/Hunsin/compass/trade"
)

var (
	pf = map[string]trade.Security{
		"2330": trade.Security{
			Market: "twse",
			Isin:   "TW0002330008",
			Symbol: "2330",
			Name:   "TSMC",
			Type:   "Stocks",
			Listed: "1994-09-05",
		},
		"0052": trade.Security{
			Market: "twse",
			Isin:   "TW0000052000",
			Symbol: "0052",
			Name:   "FB TECHNOLOGY",
			Type:   "ETF",
			Listed: "2006-09-12",
		},
		"9103": trade.Security{
			Market: "twse",
			Isin:   "TW0009103002",
			Symbol: "9103",
			Name:   "MEDTECS",
			Type:   "TDR",
			Listed: "2002-12-13",
		},
	}
)

func equal(x, y interface{}) bool {
	if z, ok := x.(trade.Quote); ok {
		q, ok := y.(trade.Quote)
		if !ok {
			return false
		}

		if math.Abs(z.Avg-q.Avg) > 0.00001 {
			return false
		}

		// assign same float numbers
		z.Avg = q.Avg
		return reflect.DeepEqual(z, q)
	}
	return reflect.DeepEqual(x, y)
}

func TestOpen(t *testing.T) {
	_, err := market.Open("twse")
	if err != nil {
		t.Fatalf("market.Open exits with error: %v", err)
	}
}
