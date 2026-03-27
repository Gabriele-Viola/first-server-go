package routes

import (
	"net/http"
	"serverGo/internal/handlers"
	"serverGo/internal/models"
)

func Configure() *http.ServeMux {
	mux := http.NewServeMux()

	mockDB := []models.User{
		{ID: 1, Name: "Gabriele", Email: "gabriele@example.com"},
		{ID: 2, Name: "Elena", Email: "elena@example.com"},
	}

	userHandler := &handlers.UserHandler{Users: mockDB}

	// "Metodo Percorso" -> Nuova sintassi super comoda
	mux.HandleFunc("GET /api/hello", handlers.GetHello)
	mux.HandleFunc("GET /api/users", userHandler.GetUsers)

	return mux
}
