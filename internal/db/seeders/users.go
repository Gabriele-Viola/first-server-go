package seeders

import (
	"database/sql"
	"fmt"
	"strings"
)

type UserSeed struct {
	Name  string
	Email string
}

func SeedUsers(tx *sql.Tx) error {
	users := []UserSeed{
		{Name: "Gabriele", Email: "gabriele@mail.com"},
		{Name: "Elena", Email: "elena@mail.com"},
	}

	if len(users) == 0 {
		return nil
	}

	values := make([]string, 0, len(users))
	args := make([]any, 0, len(users)*2)

	for _, u := range users {
		values = append(values, "(?, ?)")
		args = append(args, u.Name, u.Email)
	}

	query := `
		INSERT INTO users (name, email)
		VALUES ` + strings.Join(values, ",") + `
		ON DUPLICATE KEY UPDATE name = VALUES(name)
	`

	if _, err := tx.Exec(query, args...); err != nil {
		return fmt.Errorf("seed users: %w", err)
	}

	return nil
}
