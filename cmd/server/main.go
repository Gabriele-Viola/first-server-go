package main

import (
	"log"
	"net/http"
	"serverGo/internal/config"
	"serverGo/internal/routes"

	"github.com/joho/godotenv"

	"serverGo/internal/db"
)

func main() {

	if err := godotenv.Load(); err != nil {
		log.Println(".env non trovato")
	}

	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB config loaded for host:", cfg.DBHost)

	conn, err := db.Open(cfg.DBDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	mux := routes.Configure(conn)

	log.Println("Server avviato su :8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}

}
