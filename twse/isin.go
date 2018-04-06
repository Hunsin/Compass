package twse

import (
	"fmt"
	"io"
	"strings"

	"github.com/Hunsin/date"
	hu "github.com/Hunsin/go-htmlutil"
	"golang.org/x/net/html"
)

var (
	isin = newAgent("http://isin.twse.com.tw/isin/e_class_main.jsp?owncode=%s&market=1")
)

// parseISIN extracts the Security from urlISIN with given code
func parseISIN(code string) (*Security, error) {
	var st *Security
	var e error
	err := isin.do(func(r io.Reader) error {
		n, err := html.Parse(r)
		if err != nil {
			return err
		}

		var c *html.Node
		hu.Last(n, func(n *html.Node) (found bool) {
			if found = n.Data == "td" && hu.Text(n) == code; found {
				c = n
			}
			return
		})

		// return error if no <td> node with code found
		if c == nil {
			e = fmt.Errorf("twse: Code %s not found", code)
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
			return fmt.Errorf("twse: can not parse data of code %s", code)
		}

		d, _ := date.Parse(dateFormat, str[7])
		st = &Security{
			code: code,
			name: string(strings.TrimSpace(str[3])),
			tp:   string(strings.TrimSpace(str[5])),
			lstd: d,
		}
		return nil
	}, code)

	if e != nil {
		return st, e
	}
	return st, err
}
