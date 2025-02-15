package main

import (
	"encoding/json"
	"net/http"
)

type UUIDResponse struct {
	UUID string `json:"uuid"`
}

func uuidHandler(w http.ResponseWriter, r *http.Request) {
	uuid := generate_uuid()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(UUIDResponse{UUID: uuid})
}
