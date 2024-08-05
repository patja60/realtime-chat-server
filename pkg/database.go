package database

import (
	"database/sql"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

type DB struct {
	Conn *sql.DB
}

func New() (*DB, error) {
	connStr := "user=username dbname=mydb sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &DB{Conn: db}, nil
}
