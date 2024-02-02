package main

import (
	"fmt"
	"net/http"

	"github.com/ermapula/golang-2024/pkg/views"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/health-check", views.HealthCheck).Methods("GET")
	r.HandleFunc("/cars", views.Cars).Methods("GET")
	
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<p>/health-check - check api</p> <p>/cars - info of the race cars from Cars 2</p>")
	})
	http.ListenAndServe(":8080", r)
}
