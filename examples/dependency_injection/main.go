package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

// So we know under the covers we're ultimately using Writer to send our greeting somewhere.
// Let's use this existing abstraction to make our code testable and more reusable.
func Greet(writer io.Writer, name string) {
	fmt.Fprintf(writer, "Hello, %s", name)
}

func MyGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

// HTTP servers will be covered in a later chapter so don't worry too much about the details.
func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(MyGreeterHandler)))
}
