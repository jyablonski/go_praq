# Apuntes
[Go Guide](https://go.dev/doc/tutorial/handle-errors)

`go mod edit -replace example.com/greetings=../greetings`

## Functions
A function whose name starts with a capital letter can be called by a function not in the same package.  This is known as an exported name.

## Modules
`go mod init` creates a go.mod file to track code dependencies and their versions, giving you direct copntrol over which module versions to use.

`go mod tidy` synchronizes the module's dependencies, adding the ones required by the code.

``` sh
go mod init

# specifies that the public url path should be replaced by the local path
# to find the dependency during testing
go mod edit -replace example.com/greetings=../greetings

go mod tidy
```

## Testing
Ending a file's name with `_test` tells `go test` that this file contains test functions

``` sh
# -v lists all tests and their results
go test -v 

```

`t *testing.T` is a parameter passed into Go test functions that provides a testing context and allows you manage and report test execution details.

- It enables the use of methods like `t.Errorf()`, `t.Fatalf()`, `t.Helper()` to log test failures or mark certain functions as helper functions such as the assertion functions
- Go does not provide assertion functions out of the box, so you'll commonly see developers create their own `assertEqual` functions instead to DRY up the testing code

``` go
package main

import (
    "testing"
)

func assertEqual(t *testing.T, got, want int) {
    if got != want {
        t.Errorf("got %d, want %d", got, want)
    }
}

func TestExample(t *testing.T) {
    got := 2 + 2
    want := 4
    assertEqual(t, got, want)
}
```

## Error Handling
`error` is a built-in Go interface type to handle error values and indicate an abnormal state.

``` go
type error interface {
    Error() string
}
```

``` go
type Server struct {
	// Addr optionally specifies the TCP address for the server to listen on,
	// in the form "host:port". If empty, ":http" (port 80) is used.
	// The service names are defined in RFC 6335 and assigned by IANA.
	// See net.Dial for details of the address format.
	Addr string 
```

## Go Programs

Go Programs have a main package with a main func inside of it.

`func` defines a function with a name and body

Imports start right at the top with `import ("fmt" "testing")` type syntax

`go mod init hello` creates a module called hello in your local directory. This is required if you want to distribute your app, and it can also define dependencies'

Go Tests:

- It needs to be in a file with a name like xxx_test.go
- The test function must start with the word Test
- The test function takes one argument only t *testing.T
- To use the *testing.T type, you need to import "testing", like we did with fmt in the other file

`Errorf` prints out a message and fails the tests.  `f` here stands for format to allow you to build a string of values into the placeholder values, and it wraps those around double quotes

- `t.Errorf("got %q want %q", got, want)`
- `hello_test.go:12: got "Hello, world" want "Hello, world2"`

You can spin up Go Docs locally by running `godoc -http:8000` and going to `localhost:8000/pkg` which will have all the packages installed on your system.

- `go install golang.org/x/tools/cmd/godoc@latest`

Subtests are useful to group tests around a thing and then have subtests that describe different scenarios.  Benefit here is that shared code can be reused in other tests.

``` go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Jacob")
		want := "Hello, Jacob"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
	t.Run("say 'Hello, world' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
```

Helper Test functions can reduce code duplication and improve readability.  `t.Helper()` below will change the test failure functionality to return the line number of the actual test that failed instead of the assertCorrectMessage line number failure.

``` go
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris")
		want := "Hello, Chris"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

```

Testable Examples are documentation examples that live in the Test files to show users how to use the functions.

If Else Statements in Go are done like this

``` go
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
```

In Go, Public Functions should be capitalized and private functions should be lowercased.

Go Benchmarks allow you to test functions and run them `b.N` times to measure how long they take

- Run with `go test -bench=.`
- 80.25 ns/op means the function takes on average 80 nanoseconds to run on this computer.

``` go
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a")
	}
}

```

You can run test coverage in go by running `go test -cover`


Compile time errors are our friend because they help us write software that works,
runtime errors are our enemies because they affect our users.

Structs are a simple type in go that are just a named collection of fields where you can store data

Interfaces in Go allow you to define sets of methods without providing how the type implements them. This enables functions to accept different types as long as they implement the interface.

To implement an interface, a type must provide definitions for all the methods declared in the interface. This makes the type an instance of the interface.

You can write functions that take an interface type as a parameter. These functions can then operate on any concrete type that implements the interface, providing flexibility and promoting decoupled design.

``` go
package main

import "fmt"

// Define an interface named Speaker
// Any type implementing the Speak method, which returns a string, satisfies this interface.
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

// Function that accepts a Speaker interface type
func MakeAnimalSpeak(s Speaker) {
    fmt.Println(s.Speak())
}

func main() {
    dog := Dog{}
    cat := Cat{}

    // Both dog and cat satisfy the Speaker interface
    MakeAnimalSpeak(dog) // Outputs: Woof!
    MakeAnimalSpeak(cat) // Outputs: Meow!
}

```

- Like an `Area()` function for both Rectangle and Circle Structs

[Table driven tests](https://go.dev/wiki/TableDrivenTests) is bascially exactly like pytest mark parametrize. You can use the same code to test things but you just pass in multiple different test case inputs + their expected outputs.

The escape character prints a new line after outputting the memory address. We get the pointer (memory address) of something by placing an & character at the beginning of the symbol.

nil is synonymous with null from other programming languages. Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.

Maps allow you to store items in a manner similar to a dictionary. You can think of the key as the word and the value as the definition.  basically dictionaries in python

- Created with `dictionary map[string]string` syntax

- Should always create a dictionary like this below and initailize it to empty to avoid running into a runtime error in case it tries to write to a nil map.  both ways do the same thing

``` go
var dictionary = map[string]string{}

// OR

var dictionary = make(map[string]string)
```

Testing code that just writes print statements to stdout is pretty difficult.  Dependency Injection is helpful in these scenarios to help test your code and separate your concerns if your functions.

In `main.go` we will send to os.Stdout so our users see the countdown printed to the terminal. In tests we will send to bytes.Buffer so our tests can capture what data is being generated.


## Concurrency
Concurrency means "having more than one thing in progress." This is something that we do naturally everyday.

Normally in Go when we call a function doSomething() we wait for it to return (even if it has no value to return, we still wait for it to finish). We say that this operation is blocking - it makes us wait for it to finish. An operation that does not block in Go will run in a separate process called a goroutine. Think of a process as reading down the page of Go code from top to bottom, going 'inside' each function when it gets called to read what it does. When a separate process starts, it's like another reader begins reading inside the function, leaving the original reader to carry on going down the page.

To tell Go to start a new goroutine we turn a function call into a go statement by putting the keyword go in front of it: go doSomething().

Because the only way to start a goroutine is to put go in front of a function call, we often use anonymous functions when we want to start a goroutine. An anonymous function literal looks just the same as a normal function declaration, but without a name (unsurprisingly). You can see one above in the body of the for loop.

Anonymous functions have a number of features which make them useful, two of which we're using above. Firstly, they can be executed at the same time that they're declared - this is what the () at the end of the anonymous function is doing. Secondly they maintain access to the lexical scope in which they are defined - all the variables that are available at the point when you declare the anonymous function are also available in the body of the function.

The body of the anonymous function above is just the same as the loop body was before. The only difference is that each iteration of the loop will start a new goroutine, concurrent with the current process (the WebsiteChecker function). Each goroutine will add its result to the results map.

Sometimes, when we run our tests, two of the goroutines write to the results map at exactly the same time. Maps in Go don't like it when more than one thing tries to write to them at once, and so fatal error.

This is a race condition, a bug that occurs when the output of our software is dependent on the timing and sequence of events that we have no control over. Because we cannot control exactly when each goroutine writes to the results map, we are vulnerable to two goroutines writing to it at the same time.

- `go test -race .`

We can solve this data race by coordinating our goroutines using channels. Channels are a powerful Go data structure that can enable concurrency by using goroutines to communicate with each other and synchronize their execution. These operations, along with their details, allow communication between different processes.


## Mocking
This may or may not make the test pass for you. The problem is we're reaching out to real websites to test our own logic.

Testing code that uses HTTP is so common that Go has tools in the standard library to help you test it.

In the mocking and dependency injection chapters, we covered how ideally we don't want to be relying on external services to test our code because they can be

``` go
func Racer(a, b string) (winner string) {
	startA := time.Now()
	http.Get(a)
	aDuration := time.Since(startA)

	startB := time.Now()
	http.Get(b)
	bDuration := time.Since(startB)

	if aDuration < bDuration {
		return a
	}

	return b
}

```

## Select
`defer` is used to clean up resources, such as closing a file or in our case closing a server so that it does not continue to listen to a port. Similar to context manager in python.

You want this to execute at the end of the function, but keep the instruction near where you created the server for the benefit of future readers of the code.

Well, a chan struct{} is the smallest data type available from a memory perspective so we get no allocation versus a bool. Since we are closing and not sending anything on the chan, why allocate anything?

`ch := make(chan struct{})`

- Notice how we have to use make when creating a channel; rather than say `var ch chan struct{}`. When you use var the variable will be initialised with the "zero" value of the type. So for string it is "", int it is 0, etc.
- For channels the zero value is nil and if you try and send to it with <- it will block forever because you cannot send to nil channels

`select` allows you to wait on multiple channels. The first one to send a value "wins" and the code underneath the case is executed.

time.After is a very handy function when using select. Although it didn't happen in our case you can potentially write code that blocks forever if the channels you're listening on never return a value. `time.After` returns a chan (like ping) and will send a signal down it after the amount of time you define.

For us this is perfect; if a or b manage to return they win, but if we get to 10 seconds then our time.After will send a signal and we'll return an error.


## Reflection
You may come across scenarios though where you want to write a function where you don't know the type at compile time. Go lets us get around this with the type interface{} which you can think of as just any type (in fact, in Go any is an alias for interface{}).

- So `walk(x interface{}, fn func(string))` will accept any value for x. This is dangerous though as you lose type safety.

As a writer of such a function, you have to be able to inspect anything that has been passed to you and try and figure out what the type is and what you can do with it. This is done using reflection. This can be quite clumsy and difficult to read and is generally less performant (as you have to do checks at runtime).

- Only use this when you need to.


## Sync

In Go, a function can be associated with a type using a receiver. The receiver appears between the func keyword and the function name. It allows the function to access the fields and methods of the receiver type.

``` go
type Counter struct {
	value int
}

// c is a pointer of type Counter
// Since c is a pointer, it directly modifies the original Counter instance.
func (c *Counter) Inc() {
	c.value++
}

func (c *Counter) Value() int {
	return c.value
}

```

## Context

`context` package can be used to help debug performance problems.


## Course 1
[Link](https://www.boot.dev/courses/learn-golang)

`mySkillIssues := 42` this `:=` syntax is called the Walrus operator to initialize a new variable and assign it a value on the same line

Use named returns when documenting function signatures, they're easier to understand than not using named returns

Anonymous functions are true to form in that they have no name. They're useful when defining a function that will only be used once or to create a quick closure.

The defer keyword is a fairly unique feature of Go. It allows a function to be executed automatically just before its enclosing function returns. The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.

Deferred functions are typically used to clean up resources that are no longer being used. Often to close database connections, file handlers and the like.

Unlike Python, Go is not function-scoped, it's block-scoped. Variables declared inside a block are only accessible within that block (and its nested blocks). There's also the package scope. We'll talk about packages later, but for now, you can think of it as the outermost, nearly global scope.

Structs can be nested to represent more complex entities


``` go
type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}

myCar := car{}
myCar.frontWheel.radius = 5
```

In Go, functions that operate on specific types can be written as methods, where the function is associated with a particular type. This is why the functions area and perimeter are written with a receiver instead of starting with just the function name. 

- With methods, you can write `r.area()` instead of `area(r)`, which is cleaner and more concise.


``` go
// function without a receiver
func area(r rect) float64 {
	return r.weight * r.height
}

// you have to do `my_area = area(rect_obj)`

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	weight, height float64
}

// function with a receiver
func (r rect) area() float64 {
	return r.weight * r.height
}

// now you could do `my_area = r.area()`

type Clientable interface {
	ExampleOne(url string) ([]string, error)
	Exampletwo(image []byte) ([]float32, error)
}

type Client struct {
	*example.Client
}
```

A type satisfies an interface in Go when it implements all the methods defined by that interface.

The `go.mod` file is the module's manifest for a package. It manages the path, the Go version to use, and all dependencies

The `go.sum` file is an auto-generated file that provides checksums for all the dependencies. It ensures integrity and reproducibility of builds by verifying that the downloaded modules match the checksums.

- Automatically updated by Go tools (go build, go test, go mod tidy, etc.).
- Should never be manually edited by Developers
- Ensures the Go project will build reliably across different machines and environments

Use the `make` function when you need to allocate and initialize slices, maps, or channels. It's used specifically for types that require internal allocation and initialization before they can be used,

``` go
s := make([]int, 5)        // Creates a slice of integers with length 5, initialized with zeros
s2 := make([]int, 5, 10)   // Creates a slice with length 5 and capacity 10

m := make(map[string]int)  // Creates a map with string keys and int values
m["age"] = 30              // Add a key-value pair to the map

ch := make(chan int)       // Creates an unbuffered channel of integers
ch2 := make(chan int, 10)  // Creates a buffered channel with a capacity of 10

```


Interfaces are not classes. They are slimmer. They don't have constructors or deconstructors that require data is created or destroyed. No hierarchical nature to interfaces. They define function signatures, but not underlying behavior.

## Errors

A dangerous function is one that can potentially cause issues such as crashes, undefined behavior. 

- `panic` causes the program to terminate unless recovered. Overusing or misusing it can make your code unpredictable or lead to difficulty in debugging. Safe alternative: Return an error instead of panicking.
- `os.Exit` immediately terminates the program without calling defer statements, which can lead to resource leaks (e.g., open files or database connections).
- File/Network Operations Without Proper Cleanup, can lead to resource leaks

In Go you can see right in the function signature if a function can error or not based on if it returns an error or not. Errors are not thrown in Go, they're returned.

``` go
func getUser(user string) (string, error) {}
```

Errors are just interfaces. When something goes wrong in a function, that function should return `error` as its last return value. Any code that calls a function that can return an `error` should handle errors by testing whether the error is `nil` or not.

``` go
type error interface {
	Error() string
}

```

The errors package makes it easy to deal with errors.

- `errors.New` can be used if you just want to return an error for a specific scenario.

## Loops

Loop syntax in Go `for i := 0; i < numMessages; i++ {} etc...`

While loops don't exist in Go, you just run the for loop while only specifying a condition


``` go
// normal loop
for i := 0; i < numMessages; i++ {
	totalCost += 1.0 + (0.01 * float64(i))
}

// while loop equivalent
plantHeight := 1

for plantHeight < 5 {
	fmt.Println("still growing! current height:", plantHeight)
	plantHeight++
}

// infinite loop - bad news bears
cost := 1

for {
	cost++
}
```

## Arrays + Slices

In Go, arrays have a fixed size.

``` go
// will initialize an array of 10 zeroes
d := [10]int

// will initialize an array of 3 ints
f := [3]int{2, 3, 4}
```

Slices are built on top of arrays, which are dynamically sized and not fixed like Arrays. They're references to what's going on in an Array.

- A function that only has access to a slice can modify the underlying array

``` go
func getMessageWithRetriesForPlan(plan string) ([]string, error) {
	allMessages := getMessageWithRetries()

	if plan == planPro {
		return allMessages[:], nil // this turns an array into a slice just with [:] syntax
	}

	if plan == planFree {
		return allMessages[0:2], nil
	}

	return nil, errors.New("unsupported Plan")
}

```


``` go
// func make([]Type, length, max_cap)
// this creates a slice of 5 ints and the underlying array can contain up to 10 ints
mySlice := make([]int, 5, 10)

// capacity can also be omitted, and in this case it will default to its length
mySlice := make([]int, 5)
```

Length of a Slice is just how many elements are already allocated. Capacity is the maximum length of the slice before reallocation of the array is necessary

A variadic function can take any arbitary amount of final arguments using the `...` syntax. Similar to kwargs in python ?

`append` is a built in function used to dynamically add elements to a slice.

``` go
slice = append(slice, oneThing)
slice = append(slice, firstThing, secondThing)
slice = append(slice, anotherSlice)

// dont ever do this, you can screw up the slices and weird things can happen
someSlice = append(otherSlice, element)

// always return the result back to the same slice
otherSlice = append(otherSlice, element)
```

Can loop through slices like below:

``` go
for INDEX, ELEMENT := range SLICE {
	fmt.Println(INDEX, ELEMENT)
}
```

## Maps

Maps in Go are a data structure to hold key-value pairs, similar to dictionaries in Python. You can create them like below:

``` go
// intialize it first, then add the pairs
ages := make(map[string]int)
ages["John"] = 37
ages["Mary"] = 24
ages["Mary"] = 21

// or initialize & load it all at once
ages := map[string]int{
	"John": 37,
	"Mary": 21,
}
```

Slices, maps, and functions cannot be ucompared for equality with `==` syntax. These are essentially just pointers to memory so they aren't the same ever. For this reason, they cant be used as Map Keys when you adding key value pairs to a map.

You can however use a Struct and use it as a Key in a Map.

If you attempt to access a value from a map where the key doesn't exist, you'll return the zero value for that type. It won't panic or error out.

Functions can change the values inside a Map even if the map is created outside of the function.

A `rune` data type in Go is just a string with len(1)

## Advanced Functions

Dynamically creating functions in function signatures and passing them around as variables can seem like it adds unnecessary complexity, but there are cases where it makes sense

- HTTP API Handlers
- Pub/Sub Handlers
- Onclick callbacks
- Anytime you need to run code at a time in the future, functions as value might make sense

Anonymous functions are ones that have no name, they're often found in closures, currying, or inside higher order functions.

A first class function is a function that can be treated like any other value. `func() int` 

A higher order function is a function that accepts another function as an argument, or returns a function as a return value. `func aggregate(a, b, c int arithmetic func(int, int) int)`

Currying is the practice of writijng a function that takes a function as input, and returns a new function.

- Might be used for a Middleware in an HTTP Handler or something like that
- Kinda like injecting additional logic into some function

`defer` allows you to execute some function just before the current function exits. This is handy for:

- Closing File connections
- Closing Database connections

Closures are a function that reference variables from outside its own function body. It may access and assign to the referenced variables.

## Pointers

What do & and * do?

- & operator (address-of operator):
- It generates a pointer to its operand which gives the memory address of a variable.
- For example, z := &x means "store the memory address of x in z."


- * is used to declare a pointer type or dereference a pointer.
- & is used to get the memory address of a variable, creating a pointer.

* operator (dereference operator):

- It dereferences a pointer to gain access to the value
- It allows you to access the value stored at the memory address a pointer is pointing to.
- For example, a := *z means "get the value stored at the memory address z."
- The * operator requires a pointer as its operand so it can dereference it (retrieve the value stored at the memory address the pointer is pointing to).

When to use Pointers

- If you want a function to modify the value of a variable, you need to pass a pointer. By default, Go passes arguments by value, meaning changes inside the function won’t affect the original variable unless you use a pointer.

Nil pointers can be dangerous, if a pointer points to nothing then dereferencing it will cause a panic which crashes the program.


## Packaging

By convention, a package's naim si the same as the last element of its import path. 

- For example, the `math/rand` package would begin with `package rand`
- The `mail.io/random` package would begin with `package random`

A directory of Go code can have at most one package. All `.go` files in a difrectory must belong to the same package, or the compiler will throw an error.

- This means you only need to import code if it lives in a different directory or package.

Modules

A Go package is a collection of Go source files that are grouped together in a directory and compiled together. A package typically represents a single unit of functionality.

- A Go package contains Go code that is organized into files with a .go extension.
- Each file in a package shares the same package name (e.g., package main or package fmt).
- Packages can be imported by other Go programs to reuse the functionality provided by the package.

A Go module is a collection of Go packages that are versioned together and stored in a specific directory structure. The module system introduced in Go 1.11 allows for versioning of dependencies and provides better management of dependencies.

- A module is a collection of Go files, which can be composed of one or more packages, and it is tied to a version.
- It is defined by a `go.mod` file in the root of the module directory.
- Modules help Go handle dependency management.
- You can specify version numbers of dependencies in go.mod to ensure your code works consistently across environments.

``` go
module github.com/user/myproject

go 1.18

require (
    github.com/some/dependency v1.2.3
)
```

A Go repository is a version-controlled project (usually hosted on a platform like GitHub, GitLab, etc.) that contains Go code, and it may consist of one or more modules.

- A repository can contain multiple Go modules, and each module may contain multiple packages.
- The repository typically holds all the source code for a project and its history (commit history).
- It's a place where your Go code resides, often a Git repository hosted on a code sharing platform.

## Running Go

To run go, you can use the follow commands:

- `go run main.go`
- `go build`

`go run` will just execute your go code. This can be handy for small files but it's not preferred.

`go build` will build a compiled version of your code that can be executed on any machine, even if Go isn't installed. It's ready-to-go machine code.

- If your directory is `hello_world`, then a binary called `hello_world` will be generated by `go build`. You can run this by running `./hello_world` to run the binary
- You have to re-compile your code with `go build` if you want changes to appear
- You can also do `go build && ./hello_world` to compile and run in 1 command

Cross Compilation

- You can specify the OS, the CPU Chipset, and an optional name for the compiled binary like so:

``` sh
GOOS=darwin GOARCH=arm64 go build -o myprogram-macos-arm64
GOOS=windows GOARCH=amd64 go build -o myprogram-windows-amd64
GOOS=linux GOARCH=amd64 go build -o myprogram-linux-amd64

```

When writing functions to be used in other packages, if you want the function to be used outside of that package then you must capitalize it.

If you want to use a function from `github.com/wagslane/mystrings`, you can add it to your import list in the `*.go` file, and then it must be included in your `go.mod` file as a dependency

- It would then be usable in your `*.go` file with `mystrings.Reverse("hello world")`

Don't be exporting code or functions from the `main` package


## Concurrency & Channels

Sequential code is code that runs in order from top to bottom, 1 thing at a time.

- It's easy to read, easy to debug, and generally is used often
- However, it's not always the most performant way to run code

Concurrency is the ability to perform multiple tasks at the same time. In programming languages we're able to do this via utilizing the multiple CPU cores available on modern processors.

Concurrency is as simple as using the `go` keyword when calling a function: `go doSomething()`

- This function will be executed concurrently with the rest of the code in the function.
- The keyword is used to spawn a new goroutine, which is a lightweight thread managed by Go Runtime

Channels are a typed, thread safe queue that allow different goroutines to communicate with each other. Senders send data into the channel, and readers read from the channel. They must be created before use like so:

``` go
ch := make(chan int)

// put the value of 70 into ch
ch <- 70

// read the value of 70 into another variable
input_val := <-ch

// close the channel
close(ch)
```

- `<-` is called the channel operator. Data flows in the direction of the operator
- This operation will block until another goroutine is ready to receive the value
- Don't send a value on a closed channel because go will panic

A Deadlock is when a group of goroutines are all blocking so none of them can continue. 

Channels can be optionally buffered, where you provide a buffer length as the second argument to `make()`. 

- Sending on a buffered channel only blocks when the buffer is full

You can close channels explicitly but you don't have to, they'll still be garbage collected later on if they're finished being used.

## Mutex

A mutex in Go (short for mutual exclusion) is a synchronization primitive that provides a way to protect shared resources (e.g., variables, data structures) from being accessed by multiple goroutines at the same time. Mutexes ensure that only one goroutine can access the critical section (the code or data protected by the mutex) at any given moment, preventing race conditions.

In Go, mutexes are provided by the sync package as the sync.Mutex type.


``` go
func protected(){
	mux.Lock()
	defer mux.Unlock()
}
```

Maps are not thread safe in go, so you can't have 2 different goroutines operating on the same map. If at least one of them is writing to the map, you must lock your maps with a mutex.

## Generics

Generics are a programming feature that allows you to write flexible, reusable, and type-safe code. They let you define functions, data structures, or interfaces that can operate on different types without being tied to a specific one, while still ensuring type safety at compile time.

``` go
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

```

- `store` is a generic interface that works with any type P that satisfies the product interface.
- `product` is the constraint that ensures the type parameter P has the required methods (Price and Name).

You can define a store for any product without tying it to a specific type. For example, you could also define a Fruit type that implements product and reuse the store interface for it.

## Goose

``` sh

cd sql/schema
goose postgres postgres://postgres:postgres@localhost:5432/postgres up
```


## SQLC

SQLC generates `*.go` code for you based on provided `*.sql` files that you create to make various queries

- It relies on the `sqlc.yaml` file to find what sql files it needs and where to dump the Go code
- The generated Go code shouldn't be manually edited after it's created.
- Create the Go code w/ `sqlc generate`

# gRPC

gRPC (short for gRPC Remote Procedure Call) is a high-performance, open-source framework developed by Google that allows applications to communicate with each other, often between microservices, in a structured, efficient, and platform-neutral way. It uses Remote Procedure Calls (RPCs) to invoke methods on a remote server as if they were local, simplifying distributed application development.

- Performant Choice
- Sending JSON data w/ HTTP is expensive because of data marshalling and unmarshalling
- Uses Protocol Buffers which are a binary serialization format
- Uses HTTP/2
- Since Protobuf defines the schema, gRPC ensures type safety, enabling efficient communication and minimizing runtime errors.

Developers write a Protobuf file that describes the service, including its methods and request/response types. You then run a Protobuf Compiler command to generate the client and server stubs in the desired programming language, adn then you can implement the service logic as needed.

Have to install the following:

- [Protocol Buffer Compiler (protoc)](https://grpc.io/docs/protoc-installation/)
- [Google Go + gRPC Gen Packages](https://grpc.io/docs/languages/go/quickstart/)

``` sh

sudo apt install -y protobuf-compiler
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc \
	--proto_path=grpc_project/protobuf "grpc_project/protobuf/orders.proto" \
	--go_out=grpc_project/services/common/genproto/orders --go_opt=paths=source_relative \
	--go-grpc_out=grpc_project/services/common/genproto/orders --go-grpc_opt=paths=source_relative
```

- This command sets the file path to the protobuf/ folder and the specific protobuf file you want to compile Go code for
- You then specify the `--go_out` and `--go-grpc_out` arguments. Make sure the folder paths are already created
- Once you run it, it will generate Go code in those 2 file paths based on the Protobuf file you provide
- Do not change the file in those files

How the Workflow works for a User:

- Users make HTTP Requests to the Kitchen Service to create or view orders.
- Kitchen Service makes a gRPC call to the Orders Service.
- Orders Service returns a Protobuf response.
- Kitchen Service converts the Protobuf response to JSON.
- Kitchen Service sends the JSON response back to the user.


## Test Driven Development

1. Write a failing test
2. Write just enough code for that test to pass
3. Refactor