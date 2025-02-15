package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUUIDHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/uuid", nil)
	if err != nil {
		t.Fatalf("Could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	uuidHandler(rr, req)

	// Check the status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Check the response body
	var response UUIDResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("Could not decode response: %v", err)
	}

	fmt.Println(response)
	expected_uuid_len := 36
	uuid_len := len(response.UUID)

	if uuid_len != expected_uuid_len {
		t.Errorf("handler returned unexpected body: got %v want %v", uuid_len, expected_uuid_len)
	}
}
