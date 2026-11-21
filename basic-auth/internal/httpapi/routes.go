package httpapi

import (
	"net/http"
)

func NewRouter(h *Handler) http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/register", h.register)

	return loggingMiddleware(mux)
}
