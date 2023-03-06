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
	// Now let's make our function

	// This function returns an error
	i, err := IsText(32)
	if err != nil {
		fmt.Println("We got an error, now we handle it through code")
	}

	// This one doesn't
	i, err = IsText("Text")
	if err != nil {
		fmt.Println("We don't get an error here")
	}

	fmt.Println(i)
}

// We made a function that returns an error if its input is not text

func IsText(value any) (bool, error) {

	switch value.(type) {
	case string:
		return true, nil
	default:
		return false, &MyError{"Text is not string"}
	}

}

type MyError struct {
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("%v", e.What)
}
