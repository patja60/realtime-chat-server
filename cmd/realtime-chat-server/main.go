package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {
	fmt.Println("Hello, World!")
	// init db

	// init repo -> use db

	// init usecase -> inject repo

	// init handler -> inject usecase

	// init router -> use handler
	r := mux.NewRouter()

	r.HandleFunc("/", HelloWorldHandler)
	r.HandleFunc("/health-check", HealthCheckHandler)

	// Start the server
	log.Println("Starting server on :8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatal(err)
	}

}
