package twse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"time"
)

const urlISIN = "http://isin.twse.com.tw/isin/e_class_main.jsp?owncode=%s&market=1"

var (
	reg      = regexp.MustCompile(">.*</td>")
	tdEndTag = []byte("</td>")
	trEndTag = []byte("</tr>")
	iPermit  = make(chan bool)
)

// parseISIN extracts the Security from urlISIN with given code
func parseISIN(code string) (*Stock, error) {
	<-iPermit
	defer func() {
		time.Sleep(80 * time.Millisecond) // release after 0.08s
		go func() { iPermit <- true }()
	}()

	res, err := client.Get(fmt.Sprintf(urlISIN, code))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	out, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	tr := bytes.Split(out, trEndTag)
	if len(tr) < 2 {
		return nil, fmt.Errorf("twse: Code %s not found", code)
	}

	td := reg.FindAllSubmatch(tr[1], -1)
	if len(td) < 4 {
		return nil, fmt.Errorf("twse: Code %s not found", code)
	}

	n := bytes.TrimSuffix(td[3][0][1:], tdEndTag)
	return &Stock{
		code: code,
		name: string(bytes.TrimSpace(n)),
	}, nil
}

func init() {
	go func() { iPermit <- true }()
}
