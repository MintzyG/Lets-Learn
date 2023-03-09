package main

import "fmt"

func main() {

	// Go functions can work on multiple types if you use type parameters
	// Type parameters of a function appear between brackets before function arguments
	// 'func Index[T comparable](s []T, x T) int'
	// This declaration means that s is a slice of any type T that fulfills the built-in  constraint 'comparable'
	// 'x' is also a value of the same type

	// 'comparable' makes '==' and '!=' operators on values of the type

	// Index works for any type that supports comparison

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

}

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}
