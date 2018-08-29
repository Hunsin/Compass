package twse

import (
	"fmt"
	"io"
	"strings"

	"github.com/Hunsin/compass/trade"
	hu "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

var isin = newAgent("http://isin.twse.com.tw/isin/e_class_main.jsp?owncode=%s&market=1")

// parseISIN extracts the Security from urlISIN with given code
func parseISIN(symbol string) (trade.Security, error) {
	var s trade.Security
	var e error
	err := isin.do(func(r io.Reader) error {
		n, err := html.Parse(r)
		if err != nil {
			return err
		}

		var c *html.Node
		hu.Last(n, func(n *html.Node) (found bool) {
			if found = n.Data == "td" && hu.Text(n) == symbol; found {
				c = n
			}
			return
		})

		// return error if no <td> node with code found
		if c == nil {
			e = fmt.Errorf("twse: Code %s not found", symbol)
			return nil
		}

		// push text contents of each <td> to slice
		var str []string
		for c = c.Parent.FirstChild; c != nil; c = c.NextSibling {
			if c.Data == "td" {
				str = append(str, hu.Text(c))
			}
		}

		// prevent panic
		if len(str) < 7 {
			return fmt.Errorf("twse: can not parse data of code %s", symbol)
		}

		s = trade.Security{
			Market: "twse",
			Symbol: symbol,
			Name:   strings.TrimSpace(str[3]),
			Type:   strings.TrimSpace(str[5]),
			Listed: formatDate(str[7]),
		}
		return nil
	}, symbol)

	if e != nil {
		return s, e
	}
	return s, err
}

// An exchange implements the market.Agent interface.
type exchange struct{}

func (e *exchange) Security(symbol string) (trade.Security, error) {
	return parseISIN(symbol)
}

func (e *exchange) Listed() ([]trade.Security, error) {
	return nil, nil
}

func (e *exchange) Market() trade.Market {
	return trade.Market{
		Code:     "twse",
		Name:     "Taiwan Stock Exchange",
		Currency: "twd",
	}
}
