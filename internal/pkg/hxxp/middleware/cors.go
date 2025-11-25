package middleware

import (
	"net/http"

	"github.com/go-chi/cors"
)

func CORS(next http.Handler) http.Handler {
	opts := cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"*"},
		AllowCredentials: true,
		MaxAge:           360,
	}

	corsHandler := cors.Handler(opts)
	return corsHandler(next)
}
