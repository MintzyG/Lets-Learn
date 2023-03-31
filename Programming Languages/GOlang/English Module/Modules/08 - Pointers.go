package main

import "fmt"

// A pointer holds the memory address of a value
// The type *T is a pointer to a T value
// Its zero value is nil

func main() {
	i, j := 12, 375

	// The & operator generates a pointer to its operand
	p := &i
	// Printing the address itself
	fmt.Println(p)
	// The operator * denotes the pointer underlying value
	// Example of printing the underlying value of the pointer
	fmt.Println(*p)
	// This is known as de-referencing

	k := &j
	// You can operate through a pointer
	fmt.Println(*k)
	// Change its value
	*k = 37
	fmt.Println(*k)
	// Divide, Sum, Multiply, etc. through the pointer
	*k = *k / 37
	fmt.Println(*k)

	// Unlike C, GO has no pointer arithmetic
}
