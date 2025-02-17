package main

import (
	"context"
	"database/sql"
	"time"
)

type Reminder struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

func getReminders(db *sql.DB) ([]Reminder, error) {
	// Get today's date in YYYY-MM-DD format
	today := time.Now().Format("2006-01-02")

	// Query to find reminders for today
	query := `SELECT id, reminder FROM public.personal_reminders WHERE reminder_date = $1`

	rows, err := db.QueryContext(context.Background(), query, today)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var reminders []Reminder
	for rows.Next() {
		var r Reminder
		if err := rows.Scan(&r.ID, &r.Text); err != nil {
			return nil, err
		}
		reminders = append(reminders, r)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return reminders, nil
}
