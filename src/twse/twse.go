package twse

import (
	"crawler"
	"fmt"
	"sync"
	"time"
)

// A Security is an instance which implements crawler.Security interface.
type Security struct {
	code string
	name string
	date time.Time
}

// Symbol returns the code of a security.
func (s *Security) Symbol() string {
	return s.code
}

// Name returns the full name of a security.
func (s *Security) Name() string {
	return s.name
}

// Listed returns the date when the security listed.
func (s *Security) Listed() time.Time {
	return s.date
}

// Date returns a crawler.Daily by given date.
func (s *Security) Date(year, month, day int) (crawler.Daily, error) {
	m, err := s.Month(year, month)
	if err != nil {
		return crawler.Daily{}, err
	}

	for d := range m {
		if m[d].Date.Day() == day {
			return m[d], nil
		}
	}

	err = fmt.Errorf("twse: Given date %4d%02d%02d not fouend, is the market closed?", year, month, day)
	return crawler.Daily{}, err
}

// Month returns a list of crawler.Daily by given year and month.
func (s *Security) Month(year, month int) ([]crawler.Daily, error) {

	// return nil if s hasn't listed at the time
	if year < s.date.Year() {
		return []crawler.Daily{}, nil
	}
	if year == s.date.Year() && month < int(s.date.Month()) {
		return []crawler.Daily{}, nil
	}

	return query(s.code, year, month)
}

// Year returns a list of crawler.Daily in given year.
func (s *Security) Year(year int) ([]crawler.Daily, error) {
	m := 12
	t := time.Now()
	if year == t.Year() {
		m = int(t.Month())
	}

	ch := make(chan error)
	defer close(ch)

	wg := sync.WaitGroup{}
	yr := make(map[int][]crawler.Daily)
	for i := 0; i < m; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			var err error
			yr[j], err = s.Month(year, j+1)
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

	for i := 1; i < m; i++ {
		yr[0] = append(yr[0], yr[i]...)
	}
	return yr[0], nil
}

// An Exchange is an instance which implements crawler.Market interface.
type Exchange struct{}

// Search returns a crawler.Security by given code.
func (e *Exchange) Search(code string) (crawler.Security, error) {
	return parseISIN(code)
}

func init() {
	crawler.Register("twse", &Exchange{})
}
