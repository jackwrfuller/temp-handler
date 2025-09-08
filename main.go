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
		Handler: router,
	}

	fmt.Println("Starting server...")
	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		panic(err)
	}

}
