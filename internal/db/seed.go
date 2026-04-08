package db

import (
	"database/sql"
	"fmt"

	"serverGo/internal/db/seeders"
)

type SeederStep struct {
	Name string
	Run  func(*sql.Tx) error
}

func RunSeed(conn *sql.DB) error {
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	steps := []SeederStep{
		{Name: "users", Run: seeders.SeedUsers},
		{Name: "posts", Run: seeders.SeedPosts},
	}

	for _, s := range steps {
		if err := s.Run(tx); err != nil {
			return fmt.Errorf("seed %s: %w", s.Name, err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}
