package bucket

import (
	"database/sql"
	"fmt"
)

const (
	tableSecurities = `
		CREATE TABLE IF NOT EXISTS securities (
			id     SERIAL PRIMARY KEY,
			symbol TEXT NOT NULL,
			market TEXT NOT NULL,
			name   TEXT NOT NULL UNIQUE,
			listed DATE NOT NULL,
			type   TEXT,
			UNIQUE (symbol, market)
		);`

	tableRecords = `
		CREATE TABLE IF NOT EXISTS records (
			id       SERIAL PRIMARY KEY,
			security SERIAL NOT NULL REFERENCES securities (id),
			date     DATE   NOT NULL,
			open     DOUBLE PRECISION NOT NULL,
			high     DOUBLE PRECISION NOT NULL,
			low      DOUBLE PRECISION NOT NULL,
			close    DOUBLE PRECISION NOT NULL,
			volume   INTEGER          NOT NULL,
			avg      DOUBLE PRECISION NOT NULL,
			UNIQUE   (security, date)
		);`

	tableAverages = `
		CREATE TABLE IF NOT EXISTS averages (
			id        SERIAL PRIMARY KEY REFERENCES records (id),
			price_5   DOUBLE PRECISION NOT NULL,
			price_20  DOUBLE PRECISION,
			price_60  DOUBLE PRECISION,
			volume_5  DOUBLE PRECISION NOT NULL,
			volume_20 DOUBLE PRECISION,
			volume_60 DOUBLE PRECISION
		);`
)

// A Bucket represents a database client
type Bucket struct {
	db *sql.DB
}

// An ErrNoFound represents an error
type ErrNoFound struct {
	msg string
}

func (e *ErrNoFound) Error() string {
	return "bucket: " + e.msg
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
	for _, t := range []string{tableSecurities, tableRecords, tableAverages} {
		if _, err := b.db.Exec(t); err != nil {
			return err
		}
	}
	return nil
}
