package app

import (
	"database/sql"
	"log"
	"serverGo/internal/config"
	"serverGo/internal/db"

	"github.com/joho/godotenv"
)

func LoadConfig() (config.Config, error) {
	if err := godotenv.Load(); err != nil {
		log.Println(".env non trovato")
	}

	cfg, err := config.Load()
	if err != nil {
		return config.Config{}, err
	}
	return cfg, nil
}

func OpenDB(cfg config.Config) (*sql.DB, func(), error) {

	conn, err := db.Open(cfg.DBDSN())
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		_ = conn.Close()

	}

	return conn, cleanup, nil
}

func InitDb() (config.Config, *sql.DB, func(), error) {

	cfg, err := LoadConfig()
	if err != nil {
		return config.Config{}, nil, nil, err
	}
	conn, cleanup, err := OpenDB(cfg)
	if err != nil {
		return config.Config{}, nil, nil, err
	}
	return cfg, conn, cleanup, nil

}
