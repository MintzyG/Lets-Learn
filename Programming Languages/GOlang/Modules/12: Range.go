package main

import "fmt"

func main() {

	// The range form of the loop iterates over a slice or a map
	// When you use range two values are returned for each iteration
	// the index and a copy of the element at that index

	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Println("2 **", i, ":", v)
	}

	// You can skip getting the index or value by assigning to '_'
	// for _, v := range pow {} only returns the value
	// for i, _ := range pow {} only returns the index

	// If you only want the index you can omit the second variable
	// for i := range pow {} only returns the index

}
