package main

import "fmt"

// A var statement declares a list o variables
// It can be used both at package or function levels
// A var declaration can include initializers, one per variable
// If an initializer is present the type can be omitted, the variable will take the type of the initializer
var C, CPP, Carbon bool = true, true, true

// You can also declare multiple variables using the var()
var (
	a int     = 1
	b bool    = false
	c float64 = 3.765
)

func main() {

	fmt.Println(a, b, c)

	// Var statement used at function level without initializers
	var i, j int

	fmt.Println(C, CPP, Carbon, i, j)

	// Inside a function you can use the short assignment := in place of var with implicit type
	// Outside a function every statement must begin with a keyword(var, func, etc) so := is not available
	k := 3
	rust, python, java := true, true, "NO!"

	fmt.Println(k, rust, python, java)

	// When declaring a variable without specifying an explicit type, the type will be inferred
	// works in both var = and := expressions
	// When the right side of the declaration is typed the new variable will take its type
	//var int i
	//j := i // j is an int

	// When the right side contains an untyped numeric constant the new variable can be
	// int, float64 or complex128 depending on the needed precision
	// i := 42           | int
	// f := 3.142        | float64
	// g := 0.867 + 0.5i | complex128

	const Hello = "World!"
	// Constants are declared like variables but with the const keyword
	// Constants can be any type
	// They cannot be declared with the := syntax
	// Untyped constants are like any variable, take the type they need in that context

	// Constants as they name say, are constant, they can't have their value changed

	fmt.Println("Hello", Hello)

}
