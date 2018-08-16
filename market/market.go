package market

import (
	"errors"
	"strings"
	"sync"

	"cloud.google.com/go/civil"
	"github.com/Hunsin/compass/trade"
)

// A Security represents a financial instrument in a market.
type Security interface {
	Profile() trade.Security
	Date(year, month, date int) (trade.Quote, error)
	Month(year, month int) ([]trade.Quote, error)
	Year(int) ([]trade.Quote, error)

	// Range returns all daily quotes between the start and end date, included.
	// If the end date is before the other one, it switches them automatically.
	// Implementation of the method is optional. If not, an Err with status
	// pb.Status_UNIMPLEMENTED should returned.
	Range(start, end civil.Date) ([]trade.Quote, error)
}

// A Market represents an exchange where financial instruments are traded.
type Market interface {
	Search(symbol string) (Security, error)
	List() ([]trade.Security, error)
	Profile() *trade.Market
}

var (
	mu  sync.Mutex
	mks = make(map[string]Market)
)

// Register makes the named Market available for querying data.
func Register(name string, m Market) {
	mu.Lock()
	defer mu.Unlock()

	if m == nil {
		panic("market: A nil Market is registered")
	}

	name = strings.ToLower(name)
	if _, ok := mks[name]; ok {
		panic("market: Market " + name + " had been registered twice")
	}

	mks[name] = m
}

// Open returns a registered Market by given name.
func Open(name string) (Market, error) {
	mu.Lock()
	defer mu.Unlock()

	m, ok := mks[name]
	if !ok {
		return nil, errors.New("market: Unknown driver " + name)
	}

	return m, nil
}

// All returns all registered Markets.
func All() []Market {
	mu.Lock()
	defer mu.Unlock()

	var ms []Market
	for _, m := range mks {
		ms = append(ms, m)
	}

	return ms
}
