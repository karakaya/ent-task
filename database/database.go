package database

import (
	"context"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var DB *pgxpool.Pool

func ConnectToDB(log zerolog.Logger) error {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Info().Msg("DATABASE_URL environment variable is not set")
	}

	var err error
	// DB, err = pgxpool.New(context.Background(), connStr)
	DB, err = pgxpool.New(context.Background(), "postgres://username:password@localhost:5432/entain-task")
	if err != nil {
		log.Info().Err(err).Msg("Unable to create database connection pool")
		return err
	}

	// Test the connection
	err = DB.Ping(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Unable to ping database")
		return err
	}

	log.Info().Msg("Connected to the database successfully")
	return nil
}
