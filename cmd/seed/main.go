package main

import (
	"log"
	"serverGo/internal/config"
	"serverGo/internal/db"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println(".env non trovato")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	conn, err := db.Open(cfg.DBDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err := db.RunSeed(conn); err != nil {
		log.Fatal(err)
	}

	log.Println("seed completed")
}
