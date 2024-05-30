package main

import (
	"fmt"
)

func say_hello(name string) (string, error) {
	if name == "" {
		fmt.Printf("Wow Ok No Name\n")
	}
	fmt.Printf("Hello, %s!\n", name)
	return name, nil
}

func return_str_chars(str string) int {
	str_len := len(str)
	fmt.Println("Length of string is:", str_len)
	return str_len
}
