package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func PrimeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Primes!\n"))
}

func main() {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/api/v1", PrimeHandler)

	// Bind to a port and pass our router in
	log.Fatal(http.ListenAndServe(":8000", r))
}
