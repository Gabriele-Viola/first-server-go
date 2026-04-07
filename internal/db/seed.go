package db

import (
	"database/sql"
	"fmt"
)

func RunSeed(conn *sql.DB) error {
	tx, err := conn.Begin()
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}
	defer func() {
		_ = tx.Rollback()

	}()

	if _, err := tx.Exec(
		`INSERT INTO USERS (name, email) VALUES ('Gabriele', 'gabriele@mail.com'),('Elena', 'elena@mail.com') ON DUPLICATE KEY UPDATE name = VALUES(name)`,
	); err != nil {
		return fmt.Errorf("seed users: %w", err)
	}

	if _, err := tx.Exec(`INSERT INTO posts (user_id, title, body)
	SELECT u.id, 'Primo post', 'Contenuto primo post' FROM users u WHERE u.email = 'gabriele@mail.com'`); err != nil {
		return fmt.Errorf("seed posts gabriele: %w", err)
	}

	if _, err := tx.Exec(
		`INSERT INTO posts (user_id, title, body)
	SELECT u.id, 'Secondo post', 'Contenuto Secondo post' FROM users u WHERE u.email = 'elena@mail.com'`); err != nil {
		return fmt.Errorf("seed posts elena: %w", err)
	}
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("Commit tx: %w", err)
	}
	return nil
}
