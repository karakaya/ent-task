package migrations

import (
	"database/sql"
	"errors"
	"log"
	"os"

	"entain-golang-task/database"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgx/v5/stdlib"
)

func MigrateDB() {
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
	log.Print(os.Getwd())
	m, err := migrate.NewWithDatabaseInstance(
		"file:///app/migrations",
		"postgres", driver,
	)

	if err != nil {
		log.Fatalf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migration successful")
}
