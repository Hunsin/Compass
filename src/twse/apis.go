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
)

var (
	cst    *time.Location
	client = &http.Client{Timeout: time.Duration(8 * time.Second)}
	permit = make(chan bool)
)

type apiStock struct {
	Data   [][]string `json:"data"`
	Date   string     `json:"date"`
	Fields []string   `json:"fields"`
	Stat   string     `json:"stat"`
}

// parseAPI decodes the JSON-encoded data from given url and stores it into v
func parseAPI(url string, v interface{}) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return json.NewDecoder(res.Body).Decode(v)
}

// parseAgent calls parseAPI once permit is available and returns permit after
// it's finished. If an error occurred when calling parseAPI, it tries again
// after few seconds.
func parseAgent(url string, v interface{}) error {
	<-permit
	defer func() {
		time.Sleep(80 * time.Millisecond) // release after 0.08s
		go func() { permit <- true }()
	}()

	err := parseAPI(url, v)
	if err != nil {

		// extensive requests may cause the client being block by TWSE
		// wait for 16 seconds and try again
		time.Sleep(16 * time.Second)
		return parseAPI(url, v)
	}

	return nil
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

	st := apiStock{}
	err := parseAgent(fmt.Sprintf(urlStock, year, month, day, code), &st)
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

	go func() { permit <- true }()
}
