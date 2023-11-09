package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread of execution
// It's generally used to execute multiple functions at the same time

// Suppose we have the following function say()
func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {

	// This is how you normally run it synchronously(One func at a time)
	say("hello")

	// To invoke this function in a goroutine we use the keyword 'go' before it
	// This function will be executed concurrently
	go say("world")

	// You can also call a go routine from a anonymous function
	go func(msg string) {
		i := 0
		fmt.Println(msg)
		for {
			i++
			if i == 10 {
				fmt.Println(msg)
				break
			}
		}
	}("going")

	// Now both functions are running asynchronously in separate goroutines
	// We have to wait for them to finish before terminating the program, so we used

	time.Sleep(time.Second * 2)

	// For a better approach use WaitGroups

}
