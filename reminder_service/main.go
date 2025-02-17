package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	// get database URL from environment variable
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable is not set")
	}

	// connect to the database
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}
	defer db.Close()

	// fetch reminders
	reminders, err := getReminders(db)
	if err != nil {
		log.Fatalf("Error fetching reminders: %v", err)
	}

	// print reminders
	for _, reminder := range reminders {
		fmt.Printf("Reminder #%d: %s \n", reminder.ID, reminder.Text)
	}

	// send the slack message
	sendSlackMessage(reminders)

	fmt.Println("Sent Slack Message, exiting out ...")
}
