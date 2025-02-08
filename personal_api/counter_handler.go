package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var counter = &Counter{value: 0}

func incrementHandler(w http.ResponseWriter, r *http.Request) {
	counter.Inc()

	currentValue := counter.Value()

	response := map[string]int{
		"value": currentValue,
	}

	if currentValue >= 10 {
		response["is_lotta_requests"] = 1
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to write response: %s", err), http.StatusInternalServerError)
	}
}
