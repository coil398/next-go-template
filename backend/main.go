package main

import (
	"backend/store"

	"github.com/rs/zerolog/log"
)

func main() {
	c, err := readConfig()
	if err != nil {
		log.Error().Msgf("An error occurred while reading the env vars: %s", err)
	}

	db, err := connectToDatabase(c.DB)
	if err != nil {
		log.Fatal().Msgf("Could not connect to the database: %s", err)
		return
	}

	ds := store.NewStore(db)

	if err := startServer(c, ds); err != nil {
		log.Fatal().Msgf("Server stopped with the error: %s", err)
		return
	}
}
