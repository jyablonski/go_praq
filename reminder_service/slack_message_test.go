package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

// Test the sendSlackMessage function
func TestSendSlackMessage(t *testing.T) {
	// Set the required environment variable for testing
	os.Setenv("WEBHOOK_URL", "http://fake-url")

	// Create a mock Slack server to simulate the Slack API response
	mockSlackServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verify that the method is POST and content-type is application/json
		if r.Method != "POST" {
			t.Errorf("Expected POST request, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected Content-Type application/json, got %s", r.Header.Get("Content-Type"))
		}

		// Read the request body
		var payload map[string]interface{}
		decoder := json.NewDecoder(r.Body)
		if err := decoder.Decode(&payload); err != nil {
			t.Errorf("Failed to decode JSON body: %v", err)
		}

		// Check if the payload message is correctly formatted
		if _, ok := payload["text"]; !ok {
			t.Error("Expected 'text' field in JSON payload")
		}

		// Respond with a success HTTP status code
		w.WriteHeader(http.StatusOK)
	}))
	defer mockSlackServer.Close()

	// Override the Webhook URL with the mock server's URL
	os.Setenv("WEBHOOK_URL", mockSlackServer.URL)

	// Prepare sample reminders
	reminders := []Reminder{
		{1, "Reminder 1"},
		{2, "Reminder 2"},
	}

	// Call the sendSlackMessage function
	sendSlackMessage(reminders)
}
