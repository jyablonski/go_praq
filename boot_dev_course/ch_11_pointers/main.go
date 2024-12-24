package main

import (
	"fmt"
	"strings"
)

type Message struct {
	Recipient string
	Text      string
}

func sendMessage(m Message) {
	fmt.Printf("To: %v\n", m.Recipient)
	fmt.Printf("Message: %v\n", m.Text)
}

func removeProfanity(message *string) {
	if message == nil {
		return
	}
	messageVal := *message
	messageVal = strings.ReplaceAll(messageVal, "dang", "****")
	messageVal = strings.ReplaceAll(messageVal, "shoot", "*****")
	messageVal = strings.ReplaceAll(messageVal, "heck", "****")

	*message = messageVal
}

func Reverse(s string) string {
	res := ""

	for _, value := range s {
		// if you did this, then it adds the new character to the right
		// res = res + string(value)

		// this adds the new character to the left, effectively leaving us
		// with a reversed string
		res = string(value) + res
	}

	return res
}

func main() {
	x := 5
	y := 7

	// this will give you the memory address
	z := &x

	// this will give you the value
	a := *z

	x = 11

	fmt.Println(x, y, z, a)

	// set g to 50
	// set h equal to the memory address of g
	var g int = 50

	// set h equal to the memory address of g, which isnt an int
	// so h needs to be *int to signal it's a pointer to an int as well
	var h *int = &g

	// set the value of h to 100, which now also sets the value of g to 100
	// because they share the same memory address
	*h = 100

	fmt.Println(g, h)

	baby := Reverse("temple")
	fmt.Println(baby)
}
