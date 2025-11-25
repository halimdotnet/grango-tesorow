package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		opts := cors.Options{
			AllowedOrigins:   []string{"*"},
			AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
			AllowedHeaders:   []string{"*"},
			ExposedHeaders:   []string{"*"},
			AllowCredentials: true,
			MaxAge:           360,
		}
		cors.Handler(opts)
		next.ServeHTTP(w, r)
	})
}
