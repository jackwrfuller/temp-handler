package main

import (
	"fmt"
	"net/http"
	"strings"
	"os"

	"github.com/jackwrfuller/temp-handler/internal/controllers"
)

func main() {

	c := controllers.NewBaseHandler()

	router := http.NewServeMux()
	router.HandleFunc("/", c.HandleRequests)

	s := &http.Server{
		Addr: ":3000",
		Handler: corsMiddleware(authMiddleware(router)),
	}

	fmt.Println("Starting server...")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/api/v1/status" {
			next.ServeHTTP(w, r)
			return
		}
		expectedToken := os.Getenv("ENDPOINT_TOKEN")
		if expectedToken == "" {
			http.Error(w, "Server not configured with ENDPOINT_TOKEN", http.StatusInternalServerError)
			return
		}

		authHeader := r.Header.Get("Authorization")

		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		token := strings.TrimPrefix(authHeader, "Bearer ")
		if token != expectedToken {
			http.Error(w, "Forbidden", http.StatusForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func corsMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", "*") 
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }
        next.ServeHTTP(w, r)
    })
}
