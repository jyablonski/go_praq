package main

import (
	"errors"
	"fmt"
)

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetries() [3]string {
	return [3]string{
		"click here to sign up",
		"pretty please click here",
		"we beg you to sign up",
	}
}

func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()

	if plan == planPro {
		return allMessages[:], nil
	}

	if plan == planFree {
		return allMessages[0:2], nil
	}

	return nil, errors.New("unsupported Plan")
}

func main() {
	fmt.Println(getMessageWithRetries())
}
