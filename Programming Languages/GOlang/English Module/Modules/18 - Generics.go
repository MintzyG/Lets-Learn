package main

import "fmt"

// You can make a GO function work on multiple types using type parameters
// Type parameters of a function appear in brackets before the function's arguments
// Example:

// func Index[T comparable](s []T, x T) int {}
// Index has a 'T comparable' type parameter

// This declaration means that s is a slice of any type 'T' that fulfills the constraint 'comparable'
// x is also a value of the same type 'T'

// 'comparable' is a constraint that makes the '==' and '!=' operators available on values of the same type T
// so Index works for any type that supports comparison

// In this example Index compares a value to all slice elements until a match is found, returning that element index
// if it couldn't find any match it returns -1

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

// In addition to generic functions that support any type
// Go also supports generic types
// A type can be parameterized with a 'type parameter'
// Example:

// Here is a singly linked list that can hold any type T for example

type List[T any] struct {
	next    *List[T]
	element T
}

func main() {

	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))

}
