package main

import (
	"fmt"
	"os"
	"strconv"
)

type Config struct {
	Server *ServerConfig
	DB     *DBConfig
}

type ServerConfig struct {
	Port int
}

type DBConfig struct {
	User     string
	Password string
	Database string
	Host     string
	Port     int
	TimeZone string
}

func readDBConfig() (*DBConfig, error) {
	port, err := strconv.Atoi(os.Getenv("MYSQL_PORT"))
	if err != nil {
		return nil, fmt.Errorf("Could not read the env var 'MYSQL_PORT': %w", err)
	}

	dbConfig := &DBConfig{
		User:     os.Getenv("MYSQL_USER"),
		Password: os.Getenv("MYSQL_PASSWORD"),
		Database: os.Getenv("MYSQL_DATABASE"),
		Host:     os.Getenv("MYSQL_HOST"),
		Port:     port,
		TimeZone: os.Getenv("TZ"),
	}

	return dbConfig, nil
}

func readConfig() (*Config, error) {
	port, err := strconv.Atoi(os.Getenv("SERVER_PORT"))
	if err != nil {
		return nil, fmt.Errorf("Could not read the env var 'SERVER_PORT': %w", err)
	}

	dbConfig, err := readDBConfig()
	if err != nil {
		return nil, err
	}

	return &Config{
		&ServerConfig{
			Port: port,
		},
		dbConfig,
	}, nil
}
