package twse

import (
	"crawler"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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
	// TODO
	return nil, nil
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
