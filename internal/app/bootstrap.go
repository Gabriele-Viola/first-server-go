package app

import (
	"database/sql"
	"log"
	"serverGo/internal/config"
	"serverGo/internal/db"

	"github.com/joho/godotenv"
)

func InitDb() (config.Config, *sql.DB, func(), error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env non trovato")
	}

	cfg, err := config.Load()
	if err != nil {
		return config.Config{}, nil, nil, err
	}

	conn, err := db.Open(cfg.DBDSN())
	if err != nil {
		return config.Config{}, nil, nil, err
	}

	cleanup := func() {
		_ = conn.Close()

	}
	return cfg, conn, cleanup, nil
}
