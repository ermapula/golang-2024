package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", HealthCheck).Methods("GET")
	r.HandleFunc("/cars", Cars).Methods("GET")
	r.HandleFunc("/cars/{id}", CarsOne).Methods("GET")

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>/health-check - check api</p> <p>/cars - info of the race cars from Cars 2</p> <p>/cars/{id} - info of a specific car</p>")
	})
	http.ListenAndServe(":8080", r)
}
