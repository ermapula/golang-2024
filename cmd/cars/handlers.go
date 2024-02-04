package main

import (
	"encoding/json"
	"net/http"

	cars "github.com/ermapula/golang-2024/pkg/cars"

	"github.com/gorilla/mux"
)

type Response struct {
	Cars []cars.Car `json:"cars"`
}


func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, map[string]string{"error": message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "500 Internal Server Error")
		return
	}

	w.Header().Set("content-type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	respondWithJSON(w, http.StatusOK, map[string]string{"status": "ok", "info": cars.Info()})
}

func Cars(w http.ResponseWriter, r *http.Request) {
	cars := cars.GetCars()
	respondWithJSON(w, http.StatusOK, cars)
}

func CarsOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	car, err := cars.GetCar(id)

	if err != nil {
		respondWithError(w, http.StatusNotFound, "404 Not Found")
		return
	}

	respondWithJSON(w, http.StatusOK, car)
}
