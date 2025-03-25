package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// Registering a handler function for the /greet route
	http.HandleFunc("/greet", greet)

	// Starting the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}
