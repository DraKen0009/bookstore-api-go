package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/DraKen0009/go-bookstore/pkg/routes"
)

func main() {
	// Initialize the router
	r := mux.NewRouter()

	// Register routes for the bookstore
	routes.RegisterBookStoreRoutes(r)

	// Handle the root path with the router
	http.Handle("/", r)

	// Start the HTTP server on port 8000
	log.Println("Server started at http://localhost:8000")
	log.Fatal(http.ListenAndServe("localhost:8000", r))
}
