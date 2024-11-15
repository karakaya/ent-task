package database

import (
	"context"
	"entain-golang-task/pkg/cfg"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var DB *pgxpool.Pool

func ConnectToDB(log zerolog.Logger, config *cfg.Config) error {
	var err error
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s",
		config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Table)

	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		log.Info().Err(err).Msg("unable to create database connection pool")
		return err
	}

	err = DB.Ping(context.Background())
	if err != nil {
		log.Error().Err(err).Msg("unable to ping database")
		return err
	}

	log.Info().Msg("connected to the database successfully")

	return nil
}
