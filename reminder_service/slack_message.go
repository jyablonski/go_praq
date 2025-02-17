package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func sendSlackMessage(reminders []Reminder) {
	webhookURL := os.Getenv("WEBHOOK_URL")
	if webhookURL == "" {
		log.Println("WEBHOOK_URL environment variable is not set")
		return
	}

	var message string

	// format the reminders string we'll send over in the payload
	// by separating each reminder for the day out
	// in its own bullet point
	for _, r := range reminders {
		message += fmt.Sprintf("- %s\n", r.Text)
	}

	// creating the json payload for the slack message, which starts with Reminders:
	payload, err := json.Marshal(map[string]interface{}{"text": fmt.Sprintf("Reminders:\n%s", message)})
	if err != nil {
		log.Printf("Error marshaling JSON: %v", err)
		return
	}

	req, err := http.NewRequest("POST", webhookURL, bytes.NewBuffer(payload))
	if err != nil {
		log.Printf("Error creating request: %v", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	fmt.Println("Sending HTTP Request to Slack")
	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Error sending Slack message: %v", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Printf("Slack webhook responded with status: %d", resp.StatusCode)
	}
}
