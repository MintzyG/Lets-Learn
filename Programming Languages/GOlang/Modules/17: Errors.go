package main

import "fmt"

func main() {

	// Go expresses errors through values
	// The error type is a built-in interface similar to fmt.Stringer
	// Functions in GO often return error values and your code should handle errors by testing whether the error equals 'nil'
	// A nil error denotes success a non nil error denotes failure, and by testing errors you can handle them however you want
	// Let's create a function that returns an error if it fails
	// First we create our error struct (Look at end of code)
	// Then we make a stringer to tell the code how to print it
	// Now let's make our function then later understand what happens in 'test()'

	test()
}

// We made here a function that uses a type switch to determine if a value is of type string
// If it's not of type string it returns false and an error

func IsText(value any) (bool, error) {

	switch value.(type) {
	case string:
		return true, nil
	default:
		return false, &MyError{"Text is not string"}
	}

}

// This is our won error type, here we can specify what information our error contains

type MyError struct {
	What string
}

// This is the method that implements our error

func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}

// Now let's understand the test() function
// Here we pass a value to IsText to test if it's a string, if it's not, we return an error
// and then test to see if we got an error and handle it through our code

func test() {
	// This function returns an error since 32 is not a string
	i, err := IsText(32)
	if err != nil {
		fmt.Println(err)
		fmt.Println("As we can see, we got our error, now we handle it through code")
	}

	// This one doesn't get an error, so it passes the error check
	i, err = IsText("Text")
	if err != nil {
		fmt.Println("We don't get an error here")
	}

	fmt.Println(i)

	// If we get some very unexpected error, that we are not prepared to handle through code
	// we can throw a 'panic', that will completely stop our program and fail fast

	// We can pass strings or errors into the panic field, to display them on the console
	panic("Something went terribly wrong")

	// Panic will then exit our program with a non-zero value and show the traces of every go routine that was running on the console

}
