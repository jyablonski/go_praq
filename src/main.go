package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func Color(color string) string {
	switch color {
	case "red", "yellow", "orange":
		return "warm"
	case "blue", "green", "purple":
		return "cool"
	default:
		return "unknown"
	}
}

type GameTypesData struct {
	// Example fields based on expected JSON structure
	// Adapt this to match the JSON structure you're working with
	Field1 string `json:"field1"`
	Field2 int    `json:"field2"`
}

func QueryEndpoint(url string) (*GameTypesData, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to perform GET request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %v", err)
	}

	var data GameTypesData
	if err := json.Unmarshal(body, &data); err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON data: %v", err)
	}

	return &data, nil
}

func main() {
	df, err := QueryEndpoint("https://api.jyablonski.dev/game_types")

	fmt.Println(df.Field1, err)
}
