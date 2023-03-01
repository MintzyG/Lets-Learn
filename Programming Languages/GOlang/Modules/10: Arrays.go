package main

import "fmt"

func main() {

	// The type [n]T is an array of n elements of type T, Example:
	// var a [10]int, 'a' is an array of ten integers
	// Arrays cannot be resized since their length is part of their type
	var a [4]string
	a[0] = "Hello "
	a[1] = "my "
	a[2] = "beautiful "
	a[3] = "world!"

	// 2 methods of printing an array
	for i := 0; i < len(a); i++ {
		fmt.Printf(a[i])
	}

	fmt.Printf("\n%v", a)
}
