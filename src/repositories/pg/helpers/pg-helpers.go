package pg_helpers

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func PostgresConnect() (*sql.DB, error) {
	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%s password=%s dbname=%s sslmode=disable",
		"postgres",
		"123456789",
		"29-09-2023",
	))

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}
