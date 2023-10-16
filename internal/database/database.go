package database

import (
	"context"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/ent"
)

var db *Database

type Database struct {
	Client *ent.Client
}

func InitializeDatabase(dbHost string, dbPort string, dbUser string, dbPass string, dbName string) error {
	log.Debug().Msg("initializing database")

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPass, dbName)

	client, err := ent.Open("postgres", connectionString)
	if err != nil {
		return fmt.Errorf("failed opening connection to postgres: %v", err)
	}

	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		return fmt.Errorf("failed creating schema resources: %v", err)
	}

	db = &Database{
		Client: client,
	}

	return nil
}

func DB() *Database {
	return db
}
