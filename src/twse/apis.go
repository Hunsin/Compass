package twse

import (
	"crawler"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	urlStock   = "http://www.twse.com.tw/en/exchangeReport/STOCK_DAY?response=json&date=%4d%02d%02d&stockNo=%s"
	dateFormat = "2006/01/02"
)

var cst *time.Location

type apiStock struct {
	Data   [][]string `json:"data"`
	Date   string     `json:"date"`
	Fields []string   `json:"fields"`
	Stat   string     `json:"stat"`
}

func query(code string, year, month int) ([]crawler.Daily, error) {

	// check values
	if month < 1 || month > 12 {
		return nil, errors.New(fmt.Sprintf("twse: Invalid month %d", month))
	}
	if year < 1992 || year > time.Now().Year() {
		return nil, errors.New(fmt.Sprintf("twse: Invalid year %d", year))
	}

	// the first date available is 1992/01/04
	day := 1
	if year == 1992 {
		day = 4
	}

	res, err := http.Get(fmt.Sprintf(urlStock, year, month, day, code))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	st := apiStock{}
	err = json.NewDecoder(res.Body).Decode(&st)
	if err != nil {
		return nil, err
	}

	if st.Stat != "OK" {
		return nil, errors.New("twse: " + st.Stat)
	}

	ds := []crawler.Daily{}
	for i := range st.Data {
		t, _ := time.ParseInLocation(dateFormat, st.Data[i][0], cst)
		v, _ := strconv.Atoi(strings.Replace(st.Data[i][1], ",", "", -1))
		o, _ := strconv.ParseFloat(st.Data[i][3], 64)
		h, _ := strconv.ParseFloat(st.Data[i][4], 64)
		l, _ := strconv.ParseFloat(st.Data[i][5], 64)
		c, _ := strconv.ParseFloat(st.Data[i][6], 64)
		ds = append(ds, crawler.Daily{t, o, c, h, l, v, -1}) // Avg not support yet
	}
	return ds, nil
}

func init() {
	var err error
	cst, err = time.LoadLocation("Asia/Taipei")
	if err != nil {
		cst = time.FixedZone("Asia/Taipei", 8)
	}
}
