package main

import "fmt"

func main() {

	// The 'defer' waits until everything else in the function has returned before it executes
	// Its arguments are evaluated immediately but the action doesn't execute until the surrounding function returns

	defer fmt.Println("Hello")
	fmt.Println(" World!!")

	// Will happen before the Hello
	loop()

}

// Defers are executed in the first in last out order
func loop() {

	sum := 1
	for sum < 500 {
		sum += sum
		defer fmt.Println(sum)
	}

}
