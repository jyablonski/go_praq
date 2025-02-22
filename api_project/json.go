package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("responding with 5XX error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}
	// example struct
	// {
	// 	"error": "something went wrong"
	// }
	respondWithJSON(w, code, errResponse{Error: msg})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Println("failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)
}
