package app

import (
	"log"
	"net/http"
)

func Start() {
	mux := http.NewServeMux()

	// Registering a handler function for the /greet route
	mux.HandleFunc("/greet", greet)
	mux.HandleFunc("/customers", listCustomers)

	// Starting the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", mux))
}
