package main

import (
	"log"
	"serverGo/internal/app"
	"serverGo/internal/db"
)

func main() {
	cfg, conn, cleanup, err := app.InitDb()
	if err != nil {
		log.Fatal(err)
	}

	defer cleanup()

	if err := db.EnsureDatabase(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName); err != nil {
		log.Fatal(err)
	}

	if err := db.RunMigrations(conn); err != nil {
		log.Fatal(err)
	}

	log.Println("Migrations Completed")
}
