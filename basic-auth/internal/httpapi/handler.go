package httpapi

import (
	"encoding/json"
	"net/http"
	"time"
	"uuid"

	"github.com/egasa21/go-auth-lab/basic-auth/internal/auth"
	"github.com/egasa21/go-auth-lab/basic-auth/internal/user"
)

type Handler struct {
	store *user.Store
}

func NewHandler(store *user.Store) *Handler {
	return &Handler{store: store}
}

type RegisterResponse struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
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

	if req.Username == "" || req.Password == "" {
		http.Error(w, "username and password are required", http.StatusBadRequest)
		return
	}

	hashedPassword, err := auth.HashPassword(req.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	userReq := user.User{Username: req.Username, PasswordHash: hashedPassword}
	newUser, err := h.store.Create(ctx, userReq)
	if err != nil {
		http.Error(w, "failed to create user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(RegisterResponse{
		ID:        newUser.ID,
		Username:  newUser.Username,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	})
}
