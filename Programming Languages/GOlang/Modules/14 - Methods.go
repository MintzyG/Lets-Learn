package main

import (
	"fmt"
)

// Go doesn't have classes, but in their place GO has methods
// A method is a function with a receiver argument
// A receiver argument declares who receives that function's capabilities
// It appears between the func keyword and the function name

type Number int

type Person struct {
	Name string
	Age  int
}

// This is a method, Person receives this method
// To access this method you use the syntax of your declared person plus . and the name of the method
// Example: Maria.PersonInfo()

func (p Person) PersonInfo() {
	fmt.Println("This person is named", p.Name, "and is", p.Age, "years old")
}

func (p *Person) Birthday() {
	p.Age += 1
	fmt.Println("Happy birthday", p.Name, "you are now", p.Age, "years old")
}

// Methods can be declared in non struct types as well

func (n Number) Absolute() int {
	if n < 0 {
		return int(-n)
	}
	return int(n)
}

func main() {

	// You cannot declare a method for a type that is defined in another package

	John := Person{"John", 35}
	John.PersonInfo()

	// Remember that methods are just functions with receiver arguments
	// You can write the exact code above without a method, and it would work fine

	Negative := Number(-5)
	Positive := Number(9)

	fmt.Println(Negative.Absolute())
	fmt.Println(Positive.Absolute())

	// You can declare methods with pointer receivers
	// Methods with pointer receivers can modify the value to which the receiver points
	// Since methods often need to modify these values, pointer receivers are more common than value receivers

	// Value receivers operate on a copy of that given type
	// Pointer receivers operate on that type itself

	// John's age
	fmt.Println(John.Age)
	// Using a pointer receiver to update johns age
	John.Birthday()
	// We can see the value for John's age was permanently changed
	fmt.Println(John.Age)

	// Methods with pointer receivers can also be written as functions
	// Methods with pointer receivers can take either a value or a pointer as the receiver
	// Functions with pointer arguments can only take pointers, no values
	// The reverse is also true, functions that take values cannot take pointers, but value receivers can take both

	// In general methods on a given type should have either value or pointer receivers, not both
	// So if you need to modify a value with a pointer receiver, make all other receivers for that type pointers
	// Even if they don't need to modify a value
}
