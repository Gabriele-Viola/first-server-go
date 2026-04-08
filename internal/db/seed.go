package db

import (
	"database/sql"
	"fmt"

	"serverGo/internal/db/seeders"
)

func RunSeed(conn *sql.DB) error {
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() { _ = tx.Rollback() }()

	if err := seeders.SeedUsers(tx); err != nil {
		return err
	}
	if err := seeders.SeedPosts(tx); err != nil {
		return err
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}
	return nil
}
