package cars

import (
	"fmt"
	"net/http"
	"encoding/json"
	"strconv"

	"github.com/gorilla/mux"
)

func Cars(w http.ResponseWriter, r *http.Request) {
	var response Response
	response.Cars = cars
	
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func CarsOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idVar := vars["id"]
	id, err := strconv.Atoi(idVar)
	if err != nil {
		w.Write([]byte("Invalid id, please enter integer"))
		return
	}

	if id > len(cars) || id <= 0 {
		w.Write([]byte("A car with this id doesn't exist"))
		return
	}

	var response Response
	response.Cars = cars[id - 1: id]

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}

	w.Write(jsonResponse)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "API is up and running")
	
}
