package twse

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/text/encoding/traditionalchinese"
	"golang.org/x/text/transform"
)

const urlTable = "http://isin.twse.com.tw/isin/e_C_public.jsp?strMode=2"

// conv converts the Big5-encodeg data from given url to UTF-8
func conv(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	src := transform.NewReader(res.Body, traditionalchinese.Big5.NewDecoder())
	return ioutil.ReadAll(src)
}

// parseTable extracts the Securities from the data download from urlTable
func parseTable() ([]Stock, error) {
	out, err := conv(urlTable)
	if err != nil {
		return nil, err
	}

	frag := strings.Split(string(out), "<B>")[1:]
	last := len(frag) - 1
	frag[last] = strings.Split(frag[last], "</table>")[0]

	// TODO
	return nil, nil
}
