package main

import "fmt"

// there are interfaces which implement methods
// and there are structs which implement those interfaces
// and those structs have methods which are called by the interface methods

// Define an interface named Speaker
type Speaker interface {
	Speak() string
}

// Define a struct type 'Dog' that implements the Speaker interface
type Dog struct{}

func (d Dog) Speak() string {
	return "Woof!"
}

// Define a struct type 'Cat' that also implements the Speaker interface
type Cat struct{}

func (c Cat) Speak() string {
	return "Meow!"
}

type Bose struct{}

func (b Bose) Speak() string {
	return "Bose Speaker!"
}

// Function that accepts a Speaker interface type
func Speak(s Speaker) {
	// Type Assertion: _, ok := s.(Bose) checks if s is of type Bose.
	// The second return value (ok) is a boolean that indicates whether the assertion was successful.
	if _, ok := s.(Bose); ok {
		fmt.Println(s.Speak() + " and it's not an Animal!")
	} else {
		fmt.Println(s.Speak())
	}
}

func main() {
	dog := Dog{}
	cat := Cat{}
	bose_speaker := Bose{}
	x := 4

	Speak(dog)
	Speak(cat)
	Speak(bose_speaker)

	fmt.Println(x)

	// cannot use x (variable of type int) as Speaker value in argument to Speak: int does not implement Speaker (missing method Speak)
	// Speak(x)
}
