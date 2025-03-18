package middleware

import (
	"fmt"
	"net/http"
)

func PrintMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("Request received: %s %s\n", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
