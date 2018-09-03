package twse

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/Hunsin/compass/market"
)

var client = &http.Client{Timeout: time.Duration(8 * time.Second)}

// parse issues a GET to given URL and calls f with HTTP response body
func parse(f func(io.Reader) error, url string) error {
	res, err := client.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	return f(res.Body)
}

// An agent represents a client who is responsible to a specific URL.
// It handles all the requests to the URL and sends them periodically
// in order to avoid a mass requests block the server.
type agent struct {
	greenlight *time.Ticker
	mu         sync.Mutex
	url        string
}

// do calls parse(f, url) once it's channel is available, which url is a
// string formatted with a.url and v. If an error occurred when calling
// parse, it tries again after a few seconds.
func (a *agent) do(f func(io.Reader) error, v ...interface{}) error {
	a.mu.Lock() // make sure it's not waiting
	a.mu.Unlock()
	<-a.greenlight.C

	err := parse(f, fmt.Sprintf(a.url, v...))
	if err != nil {
		if _, ok := err.(*market.Err); !ok {
			a.mu.Lock() // block other requests
			defer a.mu.Unlock()

			// extensive requests may cause the client being block by TWSE
			// wait for 16 seconds and try again
			time.Sleep(16 * time.Second)
			err = parse(f, fmt.Sprintf(a.url, v...))
			time.Sleep(time.Second / 4)
		}
	}

	return err
}

// newAgent returns a pointer to an initialized agent with given url
func newAgent(url string) *agent {
	return &agent{
		greenlight: time.NewTicker(time.Second / 4), // release every 0.25 second
		url:        url,
	}
}
