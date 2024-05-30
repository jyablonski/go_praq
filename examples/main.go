package main

import (
	"fmt"
)

func main() {
	var input_name string
	fmt.Println("Please enter your name: ")
	fmt.Scanln(&input_name)
	say_hello(input_name)
	fmt.Println("Exiting!")
}
