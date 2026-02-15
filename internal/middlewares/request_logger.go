package middlewares

import (
	"fmt"
	"net/http"
)

func RequestLoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Request received at:", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
