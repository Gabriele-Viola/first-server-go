package models

import "time"

type User struct {
	ID    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Post struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Title     string    `json:"title"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
