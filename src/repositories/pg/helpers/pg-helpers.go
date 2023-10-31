package pg_helpers

import (
	"database/sql"
	"fmt"
	"pontos_funcionario/src/config"

	_ "github.com/lib/pq"
)

func PostgresConnect() (*sql.DB, error) {

	pgConnection, err := config.GetPGConfig()

	if err != nil {
		return nil, err
	}

	db, err := sql.Open("postgres", fmt.Sprintf(
		"user=%v password=%v dbname=%v host=%v port=%v sslmode=disable",
		pgConnection.User,
		pgConnection.Password,
		pgConnection.Database,
		pgConnection.Host,
		pgConnection.Port,
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
