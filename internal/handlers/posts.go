package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"serverGo/internal/models"
)

type PostHandler struct {
	DB *sql.DB
}

func (h *PostHandler) GetPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := h.DB.Query("SELECT id, user_id, title, body, created_at, updated_at FROM posts ORDER BY id")
	if err != nil {
		http.Error(w, "errore nel recupero post", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	posts := make([]models.Post, 0)
	for rows.Next() {
		var p models.Post
		if err := rows.Scan(&p.ID, &p.UserID, &p.Title, &p.Body, &p.CreatedAt, &p.UpdatedAt); err != nil {
			http.Error(w, "errore parsing post", http.StatusInternalServerError)
			return
		}
		posts = append(posts, p)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "errore iterazione post", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(posts)
}
