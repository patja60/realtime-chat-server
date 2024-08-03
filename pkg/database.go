package database

import "database/sql"

type DB struct {
	conn *sql.DB
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

	return &DB{conn: db}, nil
}

func (db *DB) CreateUser(email, passwordHash string) error {
	_, err := db.conn.Exec("INSERT INTO users(email, passwordHash) VALUES ($1, $2)", email, passwordHash)
	return err
}
