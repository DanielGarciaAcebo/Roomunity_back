package bootstrap

import (
	"database/sql"
	"net/http"

	// Auth
	authhttp "roomunity_back/internal/auth/adapters/http"
	authrepo "roomunity_back/internal/auth/adapters/repo"
	authapp "roomunity_back/internal/auth/app"
)

// BuildMux mount the router and make all modules
func BuildMux(db *sql.DB) *http.ServeMux {
	mux := http.NewServeMux()

	// Health / root
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			http.NotFound(w, r)
			return
		}
		_, _ = w.Write([]byte("ok\n"))
	})

	// ===== Auth module (login) =====
	usersRepo := authrepo.NewPostgresUserRepository(db)
	authService := authapp.NewAuthService(usersRepo)
	authHandler := authhttp.NewHandler(authService)
	authhttp.RegisterRoutes(mux, authHandler)

	// ===== Calendar =====

	return mux
}
