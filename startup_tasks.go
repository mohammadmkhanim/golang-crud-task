package main

import (
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func MigrateDatabase(dbURL string) {
	m, err := migrate.New(
		"file://data/migrations",
		dbURL,
	)
	if err != nil {
		log.Fatal("migration init failed:", err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		log.Fatal("migration failed:", err)
	}

	log.Println("migrations completed (or already up to date)")
}
