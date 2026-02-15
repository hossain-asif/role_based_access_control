package middlewares

import (
	"fmt"
	"net/http"
)

func JwtAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("JWT auth middleware executed.")

		next.ServeHTTP(w, r)
	})
}
