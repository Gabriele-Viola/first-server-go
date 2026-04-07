package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"serverGo/internal/models"
)

type UserHandler struct {
	DB *sql.DB
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	res := map[string]string{"status": "ok", "message": "Backend Go Standard!"}
	_ = json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	rows, err := h.DB.Query("SELECT id, name, email FROM users ORDER BY id")
	if err != nil {
		http.Error(w, "errore nel recupero utenti", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	users := make([]models.User, 0)
	for rows.Next() {
		var u models.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Email); err != nil {
			http.Error(w, "errore parsing utenti", http.StatusInternalServerError)
			return
		}
		users = append(users, u)
	}

	if err := rows.Err(); err != nil {
		http.Error(w, "errore iterazione utenti", http.StatusInternalServerError)
		return
	}

	_ = json.NewEncoder(w).Encode(users)
}
