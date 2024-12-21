package main

import (
	"errors"
	"fmt"
)

func getUser(user string) (string, error) {
	if user == "test" {
		return "", errors.New("user not found")
	}
	return user, nil
}

type userError struct {
	name string
}

func (e userError) Error() string {
	return fmt.Sprintf("%v has a problem with their account", e.name)
}

func main() {
	user, err := getUser("jacob")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("User:", user) // Output the user

	if user != "jacob" {
		fmt.Println("This isn't jacob")
		return
	}

	// profile, err := getUserPorfile(user.ID)
}
