package main

import (
	"fmt"
	"os"
	"reflect"
	"time"
)

func main() {
	// Get a greeting message and print it.
	fmt.Println("hello")

	var jacobs_var string = "tester mctesterson"

	fmt.Printf("jacobs_var is of type : %T\n", jacobs_var)

	// this caps out at 7 sig figs (LOL) and prints (3.1459265)
	const pi float32 = 3.145926535
	fmt.Println(pi)

	const pi_64 float64 = 3.145926535
	fmt.Println(pi_64)

	current_date := time.Now()
	fmt.Println(current_date)

	fmt.Println("Type of current_date:", reflect.TypeOf(current_date))

	// you cant print the type of a variable with println
	intVar := 2
	fmt.Println("intVar2 : ", intVar)
	fmt.Printf("Type of intVar2 : %T\n", intVar)

	// Check if intVar is not of type int
	// exiting with 0 means success
	// exiting with 1 means failure
	if reflect.TypeOf(intVar).Kind() != reflect.Int {
		fmt.Println("intVar is not of type int")
		os.Exit(1)
	}

	// Slice
	sliceVar := []int{2, 3, 4, 5}
	fmt.Println("sliceVar : ", sliceVar)
	fmt.Printf("Typeof sliceVar : %T\n\n", sliceVar)

	// Map
	mapVar := map[string]int{
		"one": 1,
		"two": 2,
	}
	fmt.Println("mapVar : ", mapVar)
	fmt.Printf("Typeof mapVar : %T\n\n", mapVar)

	// Struct
	type Person struct {
		Name string
		Age  int
	}
	personVar := Person{"Alice", 22}
	fmt.Println("personVar : ", personVar)
	fmt.Printf("Typeof personVar : %T\n\n", personVar)
}
