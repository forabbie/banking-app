package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
)

type Customer struct {
	Name    string `json:"name" xml:"name"`
	City    string `json:"city" xml:"city"`
	Zipcode string `json:"zipcode" xml:"zipcode"`
}

func main() {

	// Registering a handler function for the /greet route
	http.HandleFunc("/greet", greet)
	http.HandleFunc("/customers", listCustomers)

	// Starting the server on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func greet(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func listCustomers(w http.ResponseWriter, r *http.Request) {
	customers := []Customer{
		{Name: "John Doe", City: "New York", Zipcode: "10001"},
		{Name: "Jane Doe", City: "New York", Zipcode: "10001"},
	}

	if r.Header.Get("Content-Type") == "application/xml" {
		w.Header().Add("Content-Type", "application/xml")
		xml.NewEncoder(w).Encode(customers)
		return
	} else {
		w.Header().Add("Content-Type", "application/json")
		json.NewEncoder(w).Encode(customers)
	}

}
