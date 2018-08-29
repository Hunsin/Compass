package twse

import (
	"strings"

	"github.com/Hunsin/compass/market"
)

// formatDate converts date format from "2006/01/02" to "2006-01-02".
func formatDate(s string) string {
	return strings.Replace(s, "/", "-", -1)
}

// formatNum removes the "," from given string.
func formatNum(s string) string {
	return strings.Replace(s, ",", "", -1)
}

// A driver implements the market.Driver interface.
type driver struct{}

func (d *driver) Open() (market.Agent, market.Quoter, error) {
	return &exchange{}, &quoter{}, nil
}

func init() {
	market.Register("twse", &driver{})
}
