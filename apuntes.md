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

Interfaces are a powerful concept that allow you to amke functions that can be used with different types.

- Like an `Area()` function for both Rectangle and Circle Structs

[Table driven tests](https://go.dev/wiki/TableDrivenTests) is bascially exactly like pytest mark parametrize. You can use the same code to test things but you just pass in multiple different test case inputs + their expected outputs.

The escape character prints a new line after outputting the memory address. We get the pointer (memory address) of something by placing an & character at the beginning of the symbol.

nil is synonymous with null from other programming languages. Errors can be nil because the return type of Withdraw will be error, which is an interface. If you see a function that takes arguments or returns values that are interfaces, they can be nillable.