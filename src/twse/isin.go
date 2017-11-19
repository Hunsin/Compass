package twse

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"regexp"
	"time"
)

var (
	reg      = regexp.MustCompile(">.*</td>")
	tdEndTag = []byte("</td>")
	trEndTag = []byte("</tr>")
	isin     = newAgent("http://isin.twse.com.tw/isin/e_class_main.jsp?owncode=%s&market=1")
)

// parseISIN extracts the Security from urlISIN with given code
func parseISIN(code string) (*Security, error) {
	var st *Security
	var e error
	err := isin.do(func(r io.Reader) error {
		out, _ := ioutil.ReadAll(r)

		if bytes.Contains(out, []byte("an inactive ISIN")) {
			e = fmt.Errorf("twse: Code %s not found", code)
			return nil
		}

		tr := bytes.Split(out, trEndTag)
		if len(tr) < 2 {
			e = fmt.Errorf("twse: Code %s not found", code)
			return nil
		}

		td := reg.FindAllSubmatch(tr[1], -1)
		if len(td) < 8 {
			e = fmt.Errorf("twse: Code %s not found", code)
			return nil
		}

		n := bytes.TrimSuffix(td[3][0][1:], tdEndTag)
		d := bytes.TrimSuffix(td[7][0][1:], tdEndTag)
		t, _ := time.ParseInLocation(dateFormat, string(d)+" 09:00", cst)
		st = &Security{
			code: code,
			name: string(bytes.TrimSpace(n)),
			date: t,
		}
		return nil
	}, code)

	if e != nil {
		return st, e
	}
	return st, err
}
