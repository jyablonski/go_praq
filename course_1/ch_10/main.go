package main

import "fmt"

func getFormattedMessages(messages []string, formatter func(string) string) []string {
	formattedMessages := []string{}
	for _, message := range messages {
		formattedMessages = append(formattedMessages, formatter(message))
	}
	return formattedMessages
}

func getLogger(formatter func(string, string) string) func(string, string) {
	return func(a string, b string) {
		fmt.Println(formatter(a, b))
	}
}

const (
	logDeleted  = "user deleted"
	logAdmin    = "admin deleted"
	logNotFound = "user not found"
)

type user struct {
	name  string
	admin bool
}

func logAndDelete(users map[string]user, name string) (log string) {
	defer delete(users, name)
	user, ok := users[name]

	if !ok {
		return logNotFound
	}
	if user.admin {
		return logAdmin
	}

	return logDeleted
}

func adder() func(int) int {
	sum := 0

	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	fmt.Println("hello world")

}
