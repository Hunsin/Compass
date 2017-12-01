package bucket

import (
	"crawler"
	"errors"
)

const (
	querySecurity  = `SELECT * FROM securities WHERE symbol = $1;`
	insertSecurity = `INSERT INTO securities VALUES ($1, $2, $3, $4);`

	// queryLastDate = ``
)

// A Security represents a tradable financial asset which stores in a Bucket.
type Security struct {
	b      *Bucket
	id     string
	Symbol string
	Market string
	name   string
	Listed crawler.Date
	Type   string
}

// NewSecurity creates a new Security in the database by given s. If the Security
// already exists in the Bucket, it checks if the values are the same. If not, it
// returns an error.
func (b *Bucket) NewSecurity(s crawler.Security) (*Security, error) {
	if s == nil {
		return nil, errors.New("bucket: s should not be nil")
	}

	_, err := b.db.Exec(insertSecurity, s.Symbol(), s.Name(), s.Listed(), s.Type())
	if err != nil {
		return nil, err
	}

	return &Security{b: b}, nil
}

// FirstDate returns the time of the Security's earliest trading record.
// func (*Security) FirstDate() (crawler.Date, error) {

//}

// LastDate returns the time of the Security's latest trading record.
//func (*Security) LastDate() (crawler.Date, error) {

//}
