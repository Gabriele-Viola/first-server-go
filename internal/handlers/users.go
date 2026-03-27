package handlers

import (
	"encoding/json"
	"net/http"
	"serverGo/internal/models"
)

type UserHandler struct {
	Users []models.User
}

func GetHello(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Risposta semplice
	res := map[string]string{"status": "ok", "message": "Backend Go Standard!"}
	json.NewEncoder(w).Encode(res)
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Controll-Allow-Origin", "*")

	json.NewEncoder(w).Encode(h.Users)
}
