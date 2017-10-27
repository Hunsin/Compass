package crawler

import "time"

// A Daily represents the daily trading data of a Security
type Daily struct {
	Date   time.Time
	Open   float32
	Close  float32
	High   float32
	Low    float32
	Volume int
	Avg    float32
	Week   float32
	Month  float32
	Season float32
}

// A Security represents the history data of the security in a market
type Security struct {
	Symbol  string
	Name    string
	History []Daily
}
