package config

import (
	"os"
	"strconv"
)

type PgConfig struct {
	User     string
	Password string
	Port     int
	Database string
}

func GetPGConfig() (*PgConfig, error) {
	var pc PgConfig

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		pc.Port = 5432
	} else {
		pc.Port = port
	}

	pc.User = os.Getenv("DB_USER")
	pc.Password = os.Getenv("DB_PASSWORD")
	pc.Database = os.Getenv("DB_NAME")

	return &pc, nil
}
