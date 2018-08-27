package market

import (
	"time"

	"cloud.google.com/go/civil"
	"github.com/Hunsin/compass/trade"
)

// An Agent represents a market.
type Agent interface {

	// Market returns the market's information.
	Market() trade.Market

	// Listed returns all securities listed in the market.
	Listed() ([]trade.Security, error)

	// Security returns the security information with given symbol.
	Security(symbol string) (trade.Security, error)
}

// A Quoter is an instance that returns the daily trading data of securities
// in certain market.
type Quoter interface {

	// Month returns the quotes of given month.
	Month(symbol string, year int, month time.Month) ([]trade.Quote, error)

	// Year returns the quotes of given year.
	Year(symbol string, year int) ([]trade.Quote, error)

	// Range returns all daily quotes between the start and end date, included.
	// If the end date is before the other one, it switches them automatically.
	// Implementation of the method is optional. If not, an Err with status
	// pb.Status_UNIMPLEMENTED should returned.
	Range(symbol string, start, end civil.Date) ([]trade.Quote, error)
}

// A Driver is the interface that must be implemented.
type Driver interface {
	Open() (Agent, Quoter, error)
}
