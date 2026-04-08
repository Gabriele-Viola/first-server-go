package seeders

import (
	"database/sql"
	"fmt"
)

type PostSeed struct {
	UserEmail string
	Title     string
	Body      string
}

func SeedPosts(tx *sql.Tx) error {
	posts := []PostSeed{
		{
			UserEmail: "gabriele@mail.com",
			Title:     "Primo post",
			Body:      "Contenuto primo post",
		},
		{
			UserEmail: "elena@mail.com",
			Title:     "Secondo post",
			Body:      "Contenuto secondo post",
		},
	}

	for _, p := range posts {
		if _, err := tx.Exec(`
			INSERT INTO posts (user_id, title, body)
			SELECT u.id, ?, ?
			FROM users u
			WHERE u.email = ?
		`, p.Title, p.Body, p.UserEmail); err != nil {
			return fmt.Errorf("seed post (%s): %w", p.UserEmail, err)
		}
	}

	return nil
}
