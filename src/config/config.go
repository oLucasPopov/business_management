package config

import (
	"os"
	"strconv"
)

type pgConfig struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
}

type hostConfig struct {
	Host string
	Port int
}

func GetHostConfig() (*hostConfig, error) {
	var hc hostConfig
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		hc.Port = 5432
	} else {
		hc.Port = port
	}
	hc.Host = os.Getenv("SERVER_HOST")

	return &hc, nil
}

func GetPGConfig() (*pgConfig, error) {
	var pc pgConfig

	port, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		pc.Port = 5432
	} else {
		pc.Port = port
	}

	pc.Host = os.Getenv("DB_HOST")
	pc.User = os.Getenv("DB_USER")
	pc.Password = os.Getenv("DB_PASSWORD")
	pc.Database = os.Getenv("DB_NAME")

	return &pc, nil
}
