package database

import (
	"context"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/rs/zerolog/log"
	"github.com/zibbp/eos/ent"
)

var db *Database

type Database struct {
	Client *ent.Client
}

func InitializeDatabase() error {
	log.Debug().Msg("initializing database")

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

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
