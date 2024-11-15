package migrations

import (
	"database/sql"
	"ent-golang-task/database"
	"ent-golang-task/pkg/cfg"
	"errors"
	"github.com/rs/zerolog"
	"log"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/stdlib"
)

func MigrateDB(config *cfg.Config, logger zerolog.Logger) {
	connConfig := database.DB.Config().ConnConfig

	dsn := stdlib.RegisterConnConfig(connConfig)

	sqlDB, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatalf("failed to create *sql.DB from pgx.ConnConfig: %v", err)
	}
	defer sqlDB.Close()

	driver, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		log.Fatalf("failed to create migration driver: %v", err)
	}

	sourceUrl := "file://migrations"
	if config.RunningInDocker {
		sourceUrl = "file:///app/migrations"
	}
	m, err := migrate.NewWithDatabaseInstance(
		sourceUrl,
		"postgres", driver,
	)

	if err != nil {
		log.Fatalf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migration failed: %v", err)
	}

	logger.Info().Msg("migration successful")
}
