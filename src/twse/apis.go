package twse

import (
	"crawler"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	urlStock   = "http://www.twse.com.tw/en/exchangeReport/STOCK_DAY?response=json&date=%4d%02d%02d&stockNo=%s"
	dateFormat = "2006/01/02 15:04"
	limit      = 8 // limitation of parallel request
)

var (
	cst    *time.Location
	permit = make(chan bool, limit)
)

type apiStock struct {
	Data   [][]string `json:"data"`
	Date   string     `json:"date"`
	Fields []string   `json:"fields"`
	Stat   string     `json:"stat"`
}

func httpGet(url string) (*http.Response, error) {
	<-permit
	defer func() {
		time.Sleep(20 * time.Millisecond) // release after 0.02s
		permit <- true
	}()
	return http.Get(url)
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
	day := 1
	if year == 1992 {
		day = 4
	}

	res, err := httpGet(fmt.Sprintf(urlStock, year, month, day, code))
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
		return nil, fmt.Errorf("twse: %s", st.Stat)
	}

	ds := []crawler.Daily{}
	for i := range st.Data {
		// default closed time 14:30
		t, _ := time.ParseInLocation(dateFormat, st.Data[i][0]+" 14:30", cst)
		v, _ := strconv.Atoi(strings.Replace(st.Data[i][1], ",", "", -1))
		o, _ := strconv.ParseFloat(st.Data[i][3], 64)
		h, _ := strconv.ParseFloat(st.Data[i][4], 64)
		l, _ := strconv.ParseFloat(st.Data[i][5], 64)
		c, _ := strconv.ParseFloat(st.Data[i][6], 64)
		ds = append(ds, crawler.Daily{
			Date:   t,
			Open:   o,
			High:   h,
			Low:    l,
			Close:  c,
			Volume: v,
			Avg:    -1}) // Avg not support yet
	}
	return ds, nil
}

func init() {
	var err error
	cst, err = time.LoadLocation("Asia/Taipei")
	if err != nil {
		cst = time.FixedZone("Asia/Taipei", 8)
	}

	for i := 0; i < limit; i++ {
		permit <- true
	}
}
