package main

import (
	"main/cmd/db/migrate/migration"
	"main/pkg/globals"
	"main/pkg/storage"
)

func main() {
	globals.SetupEnvironmentVariables()
	storage.Connect(storage.Default())

	migration.FilterEnums()

	storage.DB.Debug().Migrator().AutoMigrate(migration.Models...)
}
