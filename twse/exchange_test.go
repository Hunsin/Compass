package twse

import (
	"errors"
	"io"
	"testing"

	hu "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

var tw = &exchange{}

func TestSecurity(t *testing.T) {
	for symbol := range pf {
		s, err := tw.Security(symbol)
		if err != nil {
			t.Fatalf("exchange.Security(%s) failed: %v", symbol, err)
		}

		if !equal(*s, pf[symbol]) {
			t.Errorf("exchange.Security(%s) failed.\nGot:  %v\nWant: %v", symbol, *s, pf[symbol])
		}
	}

	if _, err := tw.Security("0000"); err == nil {
		t.Error("exchange.Security(0000) failed. Should return not found error")
	}
}

func TestListed(t *testing.T) {
	s, err := tw.Listed()
	if err != nil {
		t.Fatal("exchange.Listed failed:", err)
	}

	var count int
	err = isin.do(func(r io.Reader) error {
		n, err := html.Parse(r)
		if err != nil {
			return err
		}

		tr := hu.Last(n, func(n *html.Node) bool {
			return n.Data == "tr"
		})

		if tr == nil {
			return errors.New("no security was found")
		}

		count, err = hu.Int(tr)
		return err
	}, "")

	if err != nil {
		t.Fatal("Could not parse ISIN page:", err)
	}

	if len(s) != count {
		t.Errorf("exchange.Listed failed: numbers got: %d; want: %d", len(s), count)
	}

	// make sure data is correct
	c := "2330"
	for i := range s {
		if s[i].Symbol == c && !equal(*s[i], pf[c]) {
			t.Errorf("exchange.Listed failed: wrong profile data.\nGot: %v\nWant: %v", s[i], pf[c])
		}
	}
}
