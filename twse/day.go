package twse

import (
	"crawler"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
	"time"

	"github.com/Hunsin/date"
)

const dateFormat = "2006/01/02"

var day = newAgent("http://www.twse.com.tw/en/exchangeReport/STOCK_DAY?response=json&date=%4d%02d%02d&stockNo=%s")

type apiDay struct {
	Data   [][]string `json:"data"`
	Date   string     `json:"date"`
	Fields []string   `json:"fields"`
	Stat   string     `json:"stat"`
}

func query(code string, year, month int) ([]crawler.Daily, error) {

	// check values
	if month < 1 || month > 12 {
		return nil, fmt.Errorf("twse: Invalid month %d", month)
	}
	if year < 1992 || year > time.Now().Year() {
		return nil, fmt.Errorf("twse: Invalid year %d", year)
	}

	// the first date available is 1992/01/04
	d := 1
	if year == 1992 {
		d = 4
	}

	st := apiDay{}
	err := day.do(func(r io.Reader) error {
		return json.NewDecoder(r).Decode(&st)
	}, year, month, d, code)
	if err != nil {
		return nil, err
	}

	if st.Stat != "OK" {
		return nil, fmt.Errorf("twse: %s", st.Stat)
	}

	ds := []crawler.Daily{}
	for i := range st.Data {
		d, _ := date.Parse(dateFormat, st.Data[i][0])
		v, _ := strconv.Atoi(strings.Replace(st.Data[i][1], ",", "", -1))
		s, _ := strconv.Atoi(strings.Replace(st.Data[i][2], ",", "", -1))
		o, _ := strconv.ParseFloat(strings.Replace(st.Data[i][3], ",", "", -1), 64)
		h, _ := strconv.ParseFloat(strings.Replace(st.Data[i][4], ",", "", -1), 64)
		l, _ := strconv.ParseFloat(strings.Replace(st.Data[i][5], ",", "", -1), 64)
		c, _ := strconv.ParseFloat(strings.Replace(st.Data[i][6], ",", "", -1), 64)
		ds = append(ds, crawler.Daily{
			Date:   d,
			Open:   o,
			High:   h,
			Low:    l,
			Close:  c,
			Volume: v,
			Avg:    float64(s) / float64(v)})
	}
	return ds, nil
}
