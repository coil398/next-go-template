package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
)

func connectToDatabase(config *DBConfig) (*sqlx.DB, error) {
	loc := strings.Replace(config.TimeZone, "/", "%2F", 1)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&%s", config.User, config.Password, config.Host, config.Port, config.Database, loc)
	log.Info().Msgf("dsn: %s", dsn)
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Info().Msgf("Connected to the database %s:%s/%s...", config.Host, config.Port, config.Database)
	return db, nil
}
