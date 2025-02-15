package main

import (
	"net/http"
	"time"
)

func healthCheckHandler(w http.ResponseWriter, _ *http.Request) {
	// You can add additional health check logic here (e.g., database, external services)

	// For this basic example, we just return HTTP 200 OK
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Sending a simple JSON response for the healthcheck
	utc_timestamp := time.Now().UTC().Format(time.RFC3339)
	w.Write([]byte(`{"status":"ok","timestamp":"` + utc_timestamp + `"}`))
}
