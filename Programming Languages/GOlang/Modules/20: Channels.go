package main

import "fmt"

// Channels are like a conduit, where you can send and receive data with the channel operator <-

// Send v to channel c with: c <- v
// Receive from channel c and assign to v with: v := <- c

// As you can see data flows in the direction of the arrow

func main() {

	// Like slices and maps channels must be created before use
	// Let's create the channel c for ints
	c := make(chan int)

	// Let's use this channel
	s := []int{1, 7, -4, 0, 3, -1}

	// Using goroutines to perform two sums async
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// retrieving the sum from the channel
	x, y := <-c, <-c

	// By default, sending and receiving lock, until the other side is ready
	// This allows goroutines to synchronize without explicit locks or conditions
	fmt.Print(x + y)

	// This code splits the task of summing up the slice into two goroutines, retrieves their sum and then finishes
	// by computing the final result
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Let's send the sum result to the channel
	c <- sum
}
