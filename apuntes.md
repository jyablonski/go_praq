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

## Error Handling
`error` is a built-in Go interface type to handle error values and indicate an abnormal state.

``` go
type error interface {
    Error() string
}
```


## What I don't like
- The forced CamelCase or MixedCaps as they call it
- The naming convention of Test Files - if it doesn't start with `Test*` then `go test` won't capture it
  - That said, this does allow for standardization of function names etc.
- This type of syntax `numbers_divisble_by_3 := []int{}` with the `[]int{}` feels bad and apparently just `[]int` works fine ?
- Print Statements are just all fucking weird.  there's `Print`, `Println`, `Errorf` like wtf ?  all we're doing is printing some output to stdout
  - They each serve their own purpose but still
- Doesn't have named parameters?  `d = Add(x=5, y=4)` ????????????????????
- It will automatically change your code when you save your file to remove things like imports.  this seems nice but wtf
- Why are dictionaries not a native data type ?
- I don't like how they do documentation.  where is the docstring? why is there like 100 lines of comments amongst 5 lines of simple code, it's just more difficult to read anything when you're just context switching

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

Interfaces in Go allow you to define sets of method signatures without providing implementations. This enables functions to accept different types as long as they implement the interface.

To implement an interface, a type must provide definitions for all the methods declared in the interface. This makes the type an instance of the interface.

You can write functions that take an interface type as a parameter. These functions can then operate on any concrete type that implements the interface, providing flexibility and promoting decoupled design.

``` go
// Sleeper interface that defines a single method `Sleep`
type Sleeper interface {
	Sleep()
}

// In Go, it's common to use empty structs when you don't need to store any data but want to define a type to implement an interface.
type DefaultSleeper struct{}

// This function implements the Sleep Method defined by the Sleepr
func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
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