package main

import (
	"log"
	"serverGo/internal/app"
	"serverGo/internal/db"
)

func main() {
	cfg, err := app.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	if err := db.EnsureDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName); err != nil {
		log.Fatal(err)
	}

	conn, cleanup, err := app.OpenDB(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer cleanup()

	if err := db.RunMigrations(conn); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations Completed")
}
