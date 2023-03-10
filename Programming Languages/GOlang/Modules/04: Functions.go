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

/*
Functions are values too, they can be passed around just like other values
Function values may be used as function arguments and return values
Example:
This function receives a function as input
this input function receives int, int as input and returns int as output
Func DoMath(fn func(int, int) int)int {}
*/

// Closures
// They are useful when you want to declare a function inline without having to name it

// This function we defined returns another function, which is anonymously defined
// the returned function closes the over the variable i to form a closure
func nextNumber() func() int {

	i := 0
	return func() int {
		i++
		return i
	}

}

func main() {
	a, b := "Hello", "World"

	fmt.Println(a, b)
	fmt.Println(add(42, 3))
	fmt.Println(swap(a, b))

	// Multiple return functions cannot have their output in a single variable context since you
	// are trying to just use one variable but are getting two
	var number1, number2 = operateXY(32)
	fmt.Println("Operate result: ", number1, number2)

	// Closures

	// We call our closure, assigning its value (a function to nextInt)
	// This function value captures its own 'i' value which will be updated each time we call nextInt
	nextInt := nextNumber()

	// Calling next int to show its 'i' value increasing
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// Calling next number to show it doesn't affect nextInt
	nextNumber()

	// nextNumber doesn't affect nextInt because it is a closure, it "captured" the value for itself
	// That's why here it still prints '4'
	fmt.Println(nextInt())

	// Variadic functions
	// Read first after the main()

	fmt.Println(1, 2, 3)

	multiply(1, 2, 3, 4)

	// If you already have multiple args in a slice, apply them in the function like this func(slice...)
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	multiply(nums...)

}

// Variadic functions
// They can be called with any number of trailing arguments for example
// Println is a variadic function since it can do this (fmt.Println(1, 2, 3))
// Let's make a function that will take an arbitrary number of int arguments
// Within the function the type 'numbers' is equivalent to []int
// So we can call len or use range on it for example
func multiply(numbers ...int) {

	var result int = 1

	for _, v := range numbers {
		result *= v
	}

	fmt.Println("The", len(numbers), "chosen multiplied resulted in:", result)

}
