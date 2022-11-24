package db_test

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func NewPostgresDB(connString string) (*sql.DB, error) {
	DB, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, fmt.Errorf("new db error %w", err)
	}

	return DB, nil
}
