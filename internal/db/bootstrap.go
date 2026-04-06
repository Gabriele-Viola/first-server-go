package db

import (
	"database/sql"
	"fmt"
)

func EnsureDatabase(host, port, user, password, dbName string) error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/?charset=utf8mb4&loc=Local", user, password, host, port)

	conn, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	defer conn.Close()

	query := fmt.Sprintf("CREATE DATABASE IF NOT EXISTS `%s` CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci", dbName)
	_, err = conn.Exec(query)
	return err
}
