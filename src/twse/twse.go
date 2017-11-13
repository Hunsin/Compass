package twse

import (
	"crawler"
	"fmt"
	"sync"
	"time"
)

var noListed = &crawler.ErrNotListed{Err: "twse: Security not listed"}

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

	// return error if s hasn't listed at the time
	if year < s.date.Year() {
		return nil, noListed
	}
	if year == s.date.Year() && month < int(s.date.Month()) {
		return nil, noListed
	}

	return query(s.code, year, month)
}

// Year returns a list of crawler.Daily in given year.
func (s *Security) Year(year int) ([]crawler.Daily, error) {
	if year < s.date.Year() {
		return nil, noListed
	}

	start := 0
	if year == s.date.Year() {
		start = int(s.date.Month()) - 1
	}

	end := 12
	t := time.Now()
	if year == t.Year() {
		end = int(t.Month())
	}

	ch := make(chan error)
	defer close(ch)

	wg := sync.WaitGroup{}
	yr := make(map[int][]crawler.Daily)
	for i := start; i < end; i++ {
		wg.Add(1)
		go func(m int) {
			defer wg.Done()
			var err error
			yr[m], err = s.Month(year, m+1)
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

// An Exchange is an instance which implements crawler.Market interface.
type Exchange struct{}

// Search returns a crawler.Security by given code.
func (e *Exchange) Search(code string) (crawler.Security, error) {
	return parseISIN(code)
}

func init() {
	crawler.Register("twse", &Exchange{})
}
