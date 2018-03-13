package twse

import (
	"fmt"
	"io"
	"net/http"
	"time"
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
	ch  chan bool
	url string
}

// do calls parse(f, url) once it's channel is available, which url is a
// string formatted with a.url and v. If an error occurred when calling
// parse, it tries again after a few seconds.
func (a *agent) do(f func(io.Reader) error, v ...interface{}) error {
	a.ch <- true
	defer func() {
		time.Sleep(100 * time.Millisecond) // release after 0.1s
		go func() { <-a.ch }()
	}()

	err := parse(f, fmt.Sprintf(a.url, v...))
	if err != nil {

		// extensive requests may cause the client being block by TWSE
		// wait for 16 seconds and try again
		time.Sleep(16 * time.Second)
		err = parse(f, fmt.Sprintf(a.url, v...))
	}

	return err
}

// newAgent returns a pointer to an initialized agent with given url
func newAgent(url string) *agent {
	return &agent{ch: make(chan bool, 1), url: url}
}
