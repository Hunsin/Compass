package twse

import (
	"crawler"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"sync"
	"time"
)

// Stock implements crawler.Security interface
type Stock struct {
	code string
	name string
}

// Symbol returns the code of a stock.
func (s *Stock) Symbol() string {
	return s.code
}

// Name returns the full name of a stock.
func (s *Stock) Name() string {
	return s.name
}

func (s *Stock) Date(year, month, date int) (crawler.Daily, error) {
	// TODO
	return crawler.Daily{}, nil
}

func (s *Stock) Month(year, month int) ([]crawler.Daily, error) {
	return query(s.code, year, month)
}

func (s *Stock) Year(year int) ([]crawler.Daily, error) {
	m := 12
	t := time.Now()
	if year == t.Year() {
		m = int(t.Month())
	}

	errs := make(chan error)
	done := make(chan bool)
	defer close(errs)
	defer close(done)

	wg := sync.WaitGroup{}
	yr := make([][]crawler.Daily, m)
	for i := 0; i < m; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			var err error
			yr[i], err = query(s.code, year, i+1)
			if err != nil {
				ch <- err
			}
		}()
	}

	go func() {
		wg.Wait()
		done <- true
	}()
	for {
		select {
		case e := <-errs:
			return nil, e
		case <-done:
			for i := 1; i < m; i++ {
				yr[0] = append(yr[0], yr[i]...)
			}
			return yr[0], nil
		default:
		}
	}
}

// Exchange implements crawler.Market interface
type Exchange struct{}

func (e *Exchange) Search(code string) (crawler.Security, error) {
	t := time.Now()
	res, err := http.Get(fmt.Sprintf(urlStock, t.Year(), t.Month(), t.Day(), code))
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
		return nil, errors.New("twse: ISIN " + code + " not found")
	}

	return &Stock{
		code: code,
		name: "", // name not support yet
	}, nil
}

func init() {
	crawler.Register("twse", &Exchange{})
}
