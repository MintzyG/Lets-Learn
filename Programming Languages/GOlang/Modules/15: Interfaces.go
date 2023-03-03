package main

import "fmt"

type MyNumber int

// An interface describes a set of methods that need to be implemented for that variable to be considered that interface type
// An interface can hold any value that implements the set of method signatures described in the interface
// Example:

// A type implements an interface by implementing its methods
// and this is an implicit implementation

type Iabsolute interface {
	Abs() MyNumber
}

func main() {

	var Numbers Iabsolute

	Number1 := MyNumber(-3)
	Numbers = Number1

	fmt.Println(Numbers.Abs())

	Number2 := MyNumber(6)
	Numbers = Number2

	fmt.Println(Numbers.Abs())

	// Under the hood interface values can be seen as a pairing of a value and a concrete type
	// (Value, Type)
	// This can be seen by printing out the interface

	var i Iabsolute
	i = MyNumber(-30)
	fmt.Printf("%v\n", i.Abs())
	Describe(i)

	// So an interface value holds a value of a specific underlying type
	// Calling a method on an interface value executes that method on its underlying type
	// as you can see in the describe function and the abs method

	// If the concrete value inside the interface is nil the method will be called with a nil receiver
	// Note that an interface value that holds a nil concrete value is itself not nil

	// A nil interface holds neither value nor concrete Type
	// Calling methods on nil interfaces results in runtime errors because there are no types inside the interface
	//to indicate which concrete method to call

	// An interface that specifies zero methods is known as an empty interface
	// An empty interface may hold values of any type
	// Since the requirement is that every type implements at least zero methods
	// Empty interfaces are used by code that handles values of unknown types for example:
	// fmt.Print, Since it can print any type, like int, string, float etc.
	// fmt.Print then takes any number of arguments of type interface{}
	var ei interface{}
	Describe2(ei)

	ei = 42
	Describe2(ei)

	ei = MyNumber(30)
	Describe2(ei)

	ei = 2.754
	Describe2(ei)

}

// This method means the type MyNumber implements the interface,
// but we don't need to explicitly declare that

func (n MyNumber) Abs() MyNumber {
	if n < 0 {
		return -n
	}
	return n
}

func Describe(i Iabsolute) {
	fmt.Printf("%v, %T \n", i, i)
}

func Describe2(i interface{}) {
	fmt.Printf("%v, %T \n", i, i)
}
