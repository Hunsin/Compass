package twse

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"time"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

const urlISIN = "http://isin.twse.com.tw/isin/e_single_main.jsp?owncode=%s"

var (
	reg      = regexp.MustCompile(">.*</td>")
	tdEndTag = []byte("</td>")
	trEndTag = []byte("</tr>")
	iPermit  = make(chan bool)
)

// conv converts the Big5-encodeg data from given url to UTF-8
func conv(url string) ([]byte, error) {
	<-iPermit
	go func() {
		time.Sleep(80 * time.Millisecond) // release after 0.08s
		iPermit <- true
	}()

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	src := transform.NewReader(res.Body, traditionalchinese.Big5.NewDecoder())
	return ioutil.ReadAll(src)
}

// parseISIN extracts the Security from urlISIN with given code
func parseISIN(code string) (*Stock, error) {
	out, err := conv(fmt.Sprintf(urlISIN, code))
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
