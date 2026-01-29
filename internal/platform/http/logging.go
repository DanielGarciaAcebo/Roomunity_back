package httpmiddleware

import (
	"log"
	"net/http"
	"time"
)

// WithRequestLogging logs method/path, remote address and duration for each request.
func WithRequestLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("HTTP %s %s from=%s took=%s", r.Method, r.URL.Path, r.RemoteAddr, time.Since(start))
	})
}
