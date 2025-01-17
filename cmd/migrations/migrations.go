package main

import (
	"authService/internal/config"
	"authService/internal/migrations"
	"authService/internal/rlog"
	"flag"
)

func main() {

	action := flag.String("action", "up", "Migration action to perform")
	flag.Parse()

	log := rlog.NewLogger("dev", "MIGRATIONS")

	if *action != "up" && *action != "down" {
		log.Fatal("Invalid migration action")
		return
	}
	cfg := config.NewConfig("dev")
	migrations.Migrate(*action, cfg)
}
