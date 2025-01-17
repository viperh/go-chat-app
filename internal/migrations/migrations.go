package migrations

import (
	"authService/internal/config"
	"authService/internal/models"
	"authService/internal/rlog"
	"authService/internal/storage/postgres"
)

var db *postgres.Postgres
var log *rlog.Logger

func Migrate(action string, cfg *config.Config) {
	log = rlog.NewLogger("dev", "MIGRATIONS")

	db = postgres.NewDatabase(cfg)

	// Migrate up
	if action == "up" {
		migrateUp()
	} else if action == "down" {
		migrateDown()
	}
}

func migrateUp() {
	err := db.Database.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatal("Could not migrate up!")
	}
}

func migrateDown() {
	err := db.Database.Migrator().DropTable(&models.User{})
	if err != nil {
		log.Fatal("Could not migrate down!")
	}
}
