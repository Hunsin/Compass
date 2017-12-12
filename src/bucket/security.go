package bucket

import (
	"crawler"
	"database/sql"
	"errors"
)

const (
	querySecurities = `SELECT * FROM securities;`
	querySecurity   = `SELECT id, name, listed, type FROM securities WHERE symbol = $1 AND market = $2;`
	insertSecurity  = `INSERT INTO securities VALUES (DEFAULT, $1, $2, $3, $4, $5) RETURNING id;`
)

// A Security represents a tradable financial asset which stores in a Bucket.
type Security struct {
	b      *Bucket
	id     string
	Symbol string
	Market string
	Name   string
	Listed crawler.Date
	Type   string
}

// Find returns a pointer to the Security with given symbol and market.
// If no Security is found, an *ErrNoFound is returned.
func (b *Bucket) Find(symbol, market string) (*Security, error) {
	sec := Security{b: b, Symbol: symbol, Market: market}
	row := b.db.QueryRow(querySecurity, symbol, market)
	err := row.Scan(&sec.id, &sec.Name, &sec.Listed, &sec.Type)
	if err != nil && err.Error() == "sql: no rows in result set" {
		return nil, &ErrNoFound{"No Security is found"}
	}
	return &sec, err
}

// NewSecurity creates a new Security in the database by given s. If the Security
// already exists in the Bucket, it checks if the values are the same. If not, it
// returns an error.
func (b *Bucket) NewSecurity(cs crawler.Security) (*Security, error) {
	if cs == nil {
		return nil, errors.New("bucket: cs should not be nil")
	}

	ns := &Security{
		b:      b,
		Symbol: cs.Symbol(),
		Market: cs.Market(),
		Name:   cs.Name(),
		Listed: cs.Listed(),
		Type:   cs.Type(),
	}

	// return the Security if it already exists;
	// return error if the properties are different
	os, err := b.Find(ns.Symbol, ns.Market)
	if err == nil {
		if ns.Name != os.Name || !ns.Listed.Equal(os.Listed) || ns.Type != os.Type {
			return nil, errors.New("bucket: The Security with different properties already exists")
		}
		return os, nil
	}

	r := b.db.QueryRow(insertSecurity, ns.Symbol, ns.Market, ns.Name, ns.Listed.String(), ns.Type)
	return ns, r.Scan(&ns.id)
}

// Securities returns all Securities with specified market stored in the Bucket.
// All Securities are returned if market equals "".
func (b *Bucket) Securities(market string) ([]*Security, error) {
	var rows *sql.Rows
	var err error
	if market == "" {
		rows, err = b.db.Query(querySecurities)
	} else {
		q := querySecurities[0:len(querySecurities)-1] + " WHERE market = $1;"
		rows, err = b.db.Query(q, market)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ss []*Security
	for rows.Next() {
		s := &Security{}
		err = rows.Scan(&s.id, &s.Symbol, &s.Market, &s.Name, &s.Listed, &s.Type)
		if err != nil {
			return nil, err
		}
		ss = append(ss, s)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	if len(ss) == 0 {
		return nil, &ErrNoFound{"No Security is found"}
	}

	return ss, nil
}
