package twse

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"sync"
	"time"

	"cloud.google.com/go/civil"
	"github.com/Hunsin/compass/market"
	"github.com/Hunsin/compass/trade"
	"github.com/Hunsin/compass/trade/pb"
)

var day = newAgent("http://www.twse.com.tw/en/exchangeReport/STOCK_DAY?response=json&date=%4d%02d%02d&stockNo=%s")

type apiDay struct {
	Data   [][]string `json:"data"`
	Date   string     `json:"date"`
	Fields []string   `json:"fields"`
	Stat   string     `json:"stat"`
}

func query(symbol string, year int, month time.Month) ([]trade.Quote, error) {

	// check values
	if month < 1 || month > 12 {
		return nil, market.Error(pb.Status_BAD_REQUEST, fmt.Sprintf("twse: Invalid month %d", month))
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
	}, year, month, d, symbol)
	if err != nil {
		return nil, err
	}

	if st.Stat != "OK" {
		return nil, fmt.Errorf("twse: %s", st.Stat)
	}

	var qs []trade.Quote
	for i := range st.Data {
		v, _ := strconv.ParseUint(formatNum(st.Data[i][1]), 10, 64) // volume
		s, _ := strconv.Atoi(formatNum(st.Data[i][2]))              // value
		o, _ := strconv.ParseFloat(formatNum(st.Data[i][3]), 64)    // open
		h, _ := strconv.ParseFloat(formatNum(st.Data[i][4]), 64)    // highest
		l, _ := strconv.ParseFloat(formatNum(st.Data[i][5]), 64)    // lowest
		c, _ := strconv.ParseFloat(formatNum(st.Data[i][6]), 64)    // close
		qs = append(qs, trade.Quote{
			Date:   formatDate(st.Data[i][0]),
			Open:   o,
			High:   h,
			Low:    l,
			Close:  c,
			Volume: v,
			Avg:    float64(s) / float64(v)})
	}
	return qs, nil
}

// A quoter is an instance which implements market.Quoter interface.
type quoter struct{}

// Month returns a list of trade.Quote by given year and month.
func (q *quoter) Month(symbol string, year int, month time.Month) ([]trade.Quote, error) {
	return query(symbol, year, month)
}

// Year returns a list of trade.Quote in given year.
func (q *quoter) Year(symbol string, year int) ([]trade.Quote, error) {
	start, end := 0, 12
	t := time.Now()
	if year == t.Year() {
		end = int(t.Month())
	}

	ch := make(chan error)
	defer close(ch)

	wg := sync.WaitGroup{}
	yr := make(map[int][]trade.Quote)
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			var err error
			yr[m], err = q.Month(symbol, year, time.Month(m+1))
			if err != nil {
				ch <- err
			}
		}(i)
	}

	go func() {
		wg.Wait()
		ch <- nil
	}()

	if err := <-ch; err != nil {
		return nil, err
	}

	for i := start + 1; i < end; i++ {
		yr[start] = append(yr[start], yr[i]...)
	}
	return yr[start], nil
}

func (q *quoter) Range(symbol string, start, end civil.Date) ([]trade.Quote, error) {
	return nil, market.Unimplemented("twse: range method not supported")
}
