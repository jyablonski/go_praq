package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIncrementHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/increment", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(incrementHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected status %v, got %v", http.StatusOK, status)
	}

	var response map[string]int
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	gotValue := response["value"]
	if gotValue != 1 {
		t.Errorf("expected value 1, got %d", gotValue)
	}

	if gotValue >= 10 {
		if _, ok := response["is_lotta_requests"]; !ok {
			t.Errorf("expected 'is_lotta_requests' key to be added")
		}
	} else {
		if _, ok := response["is_lotta_requests"]; ok {
			t.Errorf("did not expect 'is_lotta_requests' key")
		}
	}
}

func TestIncrementHandlerWithLottaRequests(t *testing.T) {
	// set the counter to 9 to check the behavior when it increments to 10
	counter.value = 9

	req, err := http.NewRequest("GET", "/increment", nil)
	if err != nil {
		t.Fatalf("could not create request: %v", err)
	}

	rr := httptest.NewRecorder()

	handler := http.HandlerFunc(incrementHandler)
	handler.ServeHTTP(rr, req)

	// check if the response contains the "is_lotta_requests" key
	var response map[string]int
	err = json.NewDecoder(rr.Body).Decode(&response)
	if err != nil {
		t.Fatalf("could not parse response: %v", err)
	}

	if _, ok := response["is_lotta_requests"]; !ok {
		t.Errorf("expected 'is_lotta_requests' key to be added when counter reaches 10")
	}
}
