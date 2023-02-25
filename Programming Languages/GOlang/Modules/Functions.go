package main

import (
	"fmt"
)

/*
A function can take zero or more arguments
When consecutive named function parameters share a type you can:
Instead of do this func add(x int, y int) omit the type from all but the last
Example: func add(x, y int) or func add(w, x, y, z int)
*/
func add(x, y int) int {
	return x + y
}

/*
A function can return any number of results
You only have to fill the return field accordingly
Example 1 return func add(x, y int) int {}
Example 2 returns func swap(a, b string) (string, string) {}
And so on
*/
func swap(a, b string) (string, string) {
	return b, a
}

/*
A function's return values may be named if they are treated as a variable inside said function
In some way declaring these variables
This should be used to document the meaning of the return values
A naked return, or a return without arguments returns the named return variables
They should not be used in longer functions as they can harm readability
*/
func operateXY(number int) (x, y int) {

	x = number * number
	y = x - number

	return
}

func main() {
	a, b := "Hello", "World"

	fmt.Println(a, b)
	fmt.Println(add(42, 3))
	fmt.Println(swap(a, b))

	// This code is not working properly
	// I have no idea why since in the go playground it works
	fmt.Println("Operate result: ", operateXY(32))
}

// Test
