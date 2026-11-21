package httpapi

import (
	"encoding/json"
	"net/http"

	"github.com/egasa21/go-auth-lab/basic-auth/internal/user"
)

type Handler struct {
	store *user.Store
}

func NewHandler(store *user.Store) *Handler {
	return &Handler{store: store}
}

func (h *Handler) register(w http.ResponseWriter, r *http.Request) {

	ctx := r.Context()

	var req struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userReq := user.User{Username: req.Username, PasswordHash: req.Password}
	if err := h.store.Create(ctx, userReq); err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("user registered successfully")
}
