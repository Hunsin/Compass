package bucket

import (
	"database/sql"
	"fmt"
)

const (
	tableSecurities = `
		CREATE TABLE IF NOT EXISTS securities (
			symbol TEXT NOT NULL PRIMARY KEY,
			name   TEXT NOT NULL,
			listed DATE NOT NULL,
			type   TEXT
		);`

	tableDaily = `
		CREATE TABLE IF NOT EXISTS daily (
			id     SERIAL PRIMARY KEY,
			symbol TEXT NOT NULL REFERENCES securities (symbol),
			date   DATE NOT NULL,
			open   DOUBLE PRECISION NOT NULL,
			high   DOUBLE PRECISION NOT NULL,
			low    DOUBLE PRECISION NOT NULL,
			close  DOUBLE PRECISION NOT NULL,
			volume INTEGER          NOT NULL,
			avg    DOUBLE PRECISION NOT NULL,
			UNIQUE (symbol, date)			
		);`

	tableAverages = `
		CREATE TABLE IF NOT EXISTS averages (
			id     SERIAL PRIMARY KEY REFERENCES daily (id),
			week   DOUBLE PRECISION NOT NULL,
			month  DOUBLE PRECISION,
			season DOUBLE PRECISION
		);`
)

// A Bucket represents a database client
type Bucket struct {
	db *sql.DB
}

// Open connects to database and initializes the db instance by given configuration
func Open(host string, port int, name, usr, pwd string, ssl bool) (*Bucket, error) {
	s := "disable"
	if ssl {
		s = "enable"
	}

	cfg := fmt.Sprintf("host=%s port=%d dbname=%s user=%s password=%s sslmode=%s",
		host, port, name, usr, pwd, s)

	var err error
	db, err := sql.Open("postgres", cfg)
	if err != nil {
		return nil, err
	}

	return &Bucket{db}, db.Ping()
}

// InitTables executes the statements which declared as table* constant
func (b *Bucket) InitTables() error {
	for _, t := range []string{tableSecurities, tableDaily, tableAverages} {
		if _, err := b.db.Exec(t); err != nil {
			return err
		}
	}
	return nil
}
