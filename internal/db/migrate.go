package db

import (
	"database/sql"
	"embed"
	"fmt"
	"sort"
	"strings"
)

//go:embed migrations/*.sql
var migrationsFS embed.FS

func RunMigrations(conn *sql.DB) error {

	entries, err := migrationsFS.ReadDir("migrations")
	if err != nil {
		return fmt.Errorf("read migrations dir: %w", err)
	}

	var files []string

	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		files = append(files, e.Name())
	}

	sort.Strings(files)

	for _, file := range files {
		content, err := migrationsFS.ReadFile("migrations/" + file)
		if err != nil {
			return fmt.Errorf("read migrations %s: %w", file, err)
		}
		stmt := strings.TrimSpace(string(content))
		if stmt == "" {
			continue
		}

		if _, err := conn.Exec(stmt); err != nil {
			return fmt.Errorf("run migration %s: %w", file, err)
		}
	}

	return nil

}
