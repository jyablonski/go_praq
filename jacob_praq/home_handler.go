package main

import (
	"encoding/json"
	"net/http"
)

// define struct for use w/ this endpoint that takes a string
// `json:"message"` ensures that the key `message` is lowercased
type HomeResponse struct {
	Message string `json:"message"`
}

// these two parameters are always used for HTTP handler functions
// ResponseWriter sends data back to the client, allows us to set headers,
// status codes, and write response bodies

// Request represents an incoming http request from a client, and contains
// url, headers, body, query params etc
func helloHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(HomeResponse{Message: "Hello, World!"})
}
