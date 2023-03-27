package main

import (
	"fmt"
	"time"
)

// Channels are like a conduit, where you can send and receive data with the channel operator <-

// Send v to channel c with: c <- v
// Receive from channel c and assign to v with: v := <- c

// Normal channels are unbuffered meaning they will only accept sends (c <- v) when there is a matching receive (v := <- c)

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
	fmt.Println(x + y)

	// This code splits the task of summing up the slice into two goroutines, retrieves their sum and then finishes
	// by computing the final result

	// Now let's say we want only a certain amount of data inside a channel
	// We can provide a buffer length for that

	// Here we created a buffered channel that can only have two variables inside it
	// If we try to put three or more values in it, we'll get an error
	c2 := make(chan string, 2)

	// With buffered channel sends only block when the buffer is full
	c2 <- "Buffered"
	c2 <- "Channel"
	// Send block

	// And receives block when the buffer is empty
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	// Receive block

	// Channel synchronization

	// We can use channels to synchronize the execution of goroutines
	// Let's use a blocking receive to wait for the goroutine to end
	// When waiting for multiple goroutines to finish you may want to use a WaitGroup

	// This is the channel that will notify the main goroutine that this process finished
	c3 := make(chan bool)
	// We start the goroutine
	go working(c3)
	// The program will only stoop running when this blocking receive gets a value from the c3 channel
	<-c3
	// Otherwise, removing this blocking receive, the program would exit even before working() started
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	// Let's send the sum result to the channel
	c <- sum
}

func working(c3 chan bool) {
	fmt.Println("Working")
	time.Sleep(time.Second)
	fmt.Println("Finished working")

	c3 <- true
}
