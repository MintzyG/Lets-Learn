package main

import "fmt"

func main() {

	// A slice is just like an array, but it doesn't have a fixed size, it is dynamically sized
	// The type []T is a slice with elements of type T
	// A slice is formed by specifying two things, a low and a high bound, separated by a colon
	// a[low : high]
	// This selects a half open range which includes the first element, but excludes the last one
	// a[1:4] this creates a slice that includes elements 1 through 3 of 'a'
	// Example:

	// An array of 6 elements
	primes := [6]int{3, 5, 7, 11, 13, 17}

	// Creates a slice that includes elements 1 through 3 of the primes array
	// Alternative declaration: slice := primes[1:4]
	var slice []int = primes[1:4]
	fmt.Println(slice)

	// A slice does not store any data it just describes a section of an underlying array
	// A slice exists through reference so changing one of its elements will change its underlying array
	// Other slices that share that underlying array will see those changes

	names := [4]string{"Sophia", "Jansen", "Gabriel", "Gustavo"}
	fmt.Println(names)

	a := names[0:2]
	b := names[0:3]
	fmt.Println(a, b)

	// We can see that changing elements in a slice changes the underlying array
	// Therefore affecting the slices as well
	b[1] = "Carlos"
	fmt.Println(a, b)
	fmt.Println(names)

	// Literals
	// A slice literal is like an array but without the length
	surnames := []string{"Machado", "Abtibol", "Leal", "Hoffmann"}
	// This creates an array, then references it

	fmt.Println(surnames)

	// Slice defaults:
	// When slicing you can omit the low or high bounds to use their defaults instead
	// The low bound default is zero, and the length of the slice is the high bound default
	slice2 := []bool{true, true, true, false, true, false, false}

	a1 := slice2[1:4] // starts at element 1 and goes up to 3
	a2 := slice2[:4]  // starts at element 0 and goes up to 3
	a3 := slice2[2:]  // Starts at element 3 and goes up to slice length
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)

	// Length and Capacity:
	// The length of a slice is the number of elements it contains
	// The capacity of a slice is the number of elements in the underlying array counting from the first element in the slice
	// Length and capacity of slices can be obtained through the functions: len(s) and cap(s)

	// This slice has a length of 6 because it has six elements
	slice3 := []int{2, 3, 5, 7, 11, 13}
	printSlice(slice3)

	// Changing the high bound to zero to give it zero length
	// Note that this does not change the underlying array, so you still have capacity 6
	slice3 = slice3[:0]
	printSlice(slice3)

	// Giving it length 4 (elements 0 through 3)
	slice3 = slice3[:4]
	printSlice(slice3)

	// Changing the low bounds of a slice drops the values that were left out if there is no declared underlying array
	// So changing the low bound of the slice will affect its capacity by dropping the first two numbers
	slice3 = slice3[2:]
	printSlice(slice3)

	// The length of a slice cannot be greater than its capacity

	// Zero value:
	// The zero value of a slice is 'nil'
	// A nil slice has length and capacity of 0 and no underlying array

	var nilSlice []int
	printSlice(nilSlice)

	if nilSlice == nil {
		fmt.Println("The slice is nil!")
	}

	// Creating a slice with make:
	// Slices can be created with the make() function, this is how you create dynamically sized arrays
	// The make function allocates a zeroed array and returns a slice that refers to that array

	//TODO: make a 'optionals' folder explaining some stuff in better detail
	makeSlice := make([]int, 5) // makes a dynamic array of length 5
	// You can pass a third argument to specify a capacity
	printSlice(makeSlice)

	// Slices of slices
	// Slices can contain any type, including other slices
	// Since I created a [][]slice there is no need to declare []int inside this slice
	board := [][]int{
		{1, 2, 3},
		{2, 3, 4},
		{3, 4, 5},
	}
	// Board is a matrix of 3x3 proportions, but you are not limited to a fixed size like that
	// You can have for example a slice of cap = 3 len = 3 that contains the following slices
	// []int{1, 2, 3, 4, ... , 99}, []int{1}, []int{1, 2, 3, 4, 5}

	fmt.Println(board[2][2])
	fmt.Println(board[1][2])

	// Appending
	// Lets say you created a nil slice, but you want elements in it
	// You can add them by using the append() function
	var nilSlice2 []int
	printSlice(nilSlice2)
	// This function is appending the values 1, 2 and 3 to the nil slice
	// The first parameter is the slice you'll be appending to, the rest are the values you are appending
	nilSlice2 = append(nilSlice2, 1, 2, 3)
	printSlice(nilSlice2)
	// If the slice you are appending to doesn't have space for more values, the append function will
	// allocate a bigger array, the returned slice will point to this new array

}

func printSlice(slice []int) {
	fmt.Println("len = ", len(slice), " cap = ", cap(slice), " slice = ", slice)
}
