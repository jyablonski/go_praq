package main

import (
	"fmt"
	"time"
)

type email struct {
	date time.Time
}

func sendEmail(message string) {

	// this is an anonymous function
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func filterOldEmails(emails []email) {
	isOldChan := make(chan bool)

	// wrap it in an anonymous function goroutine so it starts executing and then moves on
	go func() {
		for _, e := range emails {
			if e.date.Before(time.Date(2020, 0, 0, 0, 0, 0, 0, time.UTC)) {
				isOldChan <- true
				continue
			}
			isOldChan <- false
		}
	}()

	// this will run in a separate goroutine whenever it's ready
	isOld := <-isOldChan
	fmt.Println("email 1 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 2 is old:", isOld)
	isOld = <-isOldChan
	fmt.Println("email 3 is old:", isOld)
}

func waitForDbs(numDBs int, dbChan chan struct{}) {

	for i := 0; i < numDBs; i++ {
		<-dbChan
	}

}

func main() {
	sendEmail("jacobs world")

	// Example usage of filterOldEmails
	emails := []email{
		{date: time.Date(2019, time.December, 31, 23, 59, 59, 0, time.UTC)},
		{date: time.Date(2021, time.January, 1, 0, 0, 0, 0, time.UTC)},
		{date: time.Date(2018, time.March, 15, 0, 0, 0, 0, time.UTC)},
	}

	filterOldEmails(emails)
}
