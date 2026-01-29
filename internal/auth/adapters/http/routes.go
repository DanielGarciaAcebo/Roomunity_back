package http

import "net/http"

// RegisterRoutes auth routes
func RegisterRoutes(mux *http.ServeMux, h *Handler) {
	mux.HandleFunc("/auth/login", h.Login)
	//mux.HandleFunc("/auth/register", h.register)
}
