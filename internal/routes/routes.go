package routes

import (
	"database/sql"
	"net/http"
	"serverGo/internal/handlers"
)

func Configure(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	userHandler := &handlers.UserHandler{DB: db}
	postHandler := &handlers.PostHandler{DB: db}

	mux.HandleFunc("GET /api/hello", handlers.GetHello)
	mux.HandleFunc("GET /api/users", userHandler.GetUsers)
	mux.HandleFunc("GET /api/posts", postHandler.GetPosts)

	return mux
}
