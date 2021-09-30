package repository

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresDB() (*sql.DB, error) {
	db, err := sql.Open("postgres",
		fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
			host, port, username, dbname, password, sslmode))
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
