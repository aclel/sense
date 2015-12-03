package store

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

// Custom type which wraps the db connection pool.
// This makes it possible to mock the database for unit testing.
type DB struct {
	*sqlx.DB
}

// Opens a database with the given data source name.
// This does not create a connection, it just validates
// the arguments and pings the database to ensure a connection
// is possible.
//
// This method should be called just once.
//
// Returns a *sqlx.DB which is a pointer to a database handle
// that maintains a pool of underlying connections. It's safe
// for concurrent use by multiple go routines.
func NewDB(dataSourceName string) (*DB, error) {
	var err error
	db, err := sqlx.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &DB{db}, nil
}
