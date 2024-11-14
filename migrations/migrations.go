package migrations

import (
	"database/sql"
	"log"

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

	m, err := migrate.NewWithDatabaseInstance(
		"file://migrations",
		"postgres", driver,
	)
	if err != nil {
		log.Fatalf("failed to initialize migrate instance: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	}

	log.Println("migration successful")
}
