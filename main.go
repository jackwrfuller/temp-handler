package main

import (
	"fmt"
	"net/http"

	"github.com/jackwrfuller/temp-handler/internal/controllers"
)

func main() {

	c := controllers.NewBaseHandler()

	router := http.NewServeMux()
	router.HandleFunc("/", c.HandleRequests)

	s := &http.Server{
		Addr: ":3000",
		Handler: corsMiddleware(router),
	}

	fmt.Println("Starting server...")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

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
