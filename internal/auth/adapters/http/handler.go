package http

import (
	"encoding/json"
	"log"
	"net/http"

	authapp "roomunity_back/internal/auth/app"
)

// Handler translates HTTP requests/responses to application use-cases.
type Handler struct {
	service *authapp.AuthService
}

// NewHandler builds a new HTTP handler for auth endpoints.
func NewHandler(service *authapp.AuthService) *Handler {
	return &Handler{service: service}
}

type credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Token string       `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID        int64  `json:"id"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Group     string `json:"group"`
	Number    int    `json:"number"`
}

// Login handles POST /auth/login.
// NOTE: Password is compared in plaintext for now (dev only). Register will introduce hashing later.
func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	// Allow CORS preflight requests.
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var cred credentials
	if err := json.NewDecoder(r.Body).Decode(&cred); err != nil {
		http.Error(w, "invalid JSON body", http.StatusBadRequest)
		return
	}

	// Log attempt without printing the password.
	log.Printf("AUTH LOGIN attempt username=%q from=%s", cred.Username, r.RemoteAddr)

	u, err := h.service.Login(cred.Username, cred.Password)
	if err != nil {
		log.Printf("AUTH LOGIN error username=%q err=%v", cred.Username, err)
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	if u == nil {
		log.Printf("AUTH LOGIN failed username=%q", cred.Username)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusUnauthorized)
		_ = json.NewEncoder(w).Encode(map[string]string{"message": "invalid credentials"})
		return
	}

	log.Printf("AUTH LOGIN success username=%q id=%d", cred.Username, u.ID)

	// Token is hardcoded for now (dev only). We'll replace it with JWT later.
	resp := AuthResponse{
		Token: "DEV_TOKEN_123",
		User: UserResponse{
			ID:        u.ID,
			Email:     u.Email,
			Username:  u.Username,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Group:     u.GroupName,
			Number:    u.Number,
		},
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(resp)
}
