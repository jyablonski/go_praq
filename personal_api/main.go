package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", HomeHandler) // Route requests to the handler function
	http.HandleFunc("/increment", incrementHandler)
	http.ListenAndServe(":8080", nil) // Start server on port 8080
}
