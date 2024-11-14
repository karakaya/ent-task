package database

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var DB *pgxpool.Pool

func ConnectToDB(log zerolog.Logger) error {
	var err error
	DB, err = pgxpool.New(context.Background(), "postgres://username:password@localhost:5432/entain-task")
	if err != nil {
		log.Info().Err(err).Msg("Unable to create database connection pool")
		return err
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("Unable to ping database")
		return err
	}

	log.Info().Msg("Connected to the database successfully")

	return nil
}
