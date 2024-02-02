package views

import (
	"fmt"
	"net/http"
	"encoding/json"
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

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	fmt.Fprintf(w, "API is up and running")
	
}
