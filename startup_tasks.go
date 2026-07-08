package main

import (
	"TaskCrud/utils"

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
		utils.LogError("MigrateDatabase", "migration init failed: {0}", err)
		panic(err)
	}

	err = m.Up()

	if err != nil && err != migrate.ErrNoChange {
		utils.LogError("MigrateDatabase", "migration failed: {0}", err)
		panic(err)
	}

	utils.LogSuccess("MigrateDatabase", "migrations completed (or already up to date)")
}
