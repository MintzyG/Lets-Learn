package main

import (
	"fmt"
)

// A sender can close a channel to indicate no more values will be sent
// Receivers can test is a channel has been closed by assigning a second parameter to the receiving expression
// v, ok := <- c
// receives the value and checks if the channel has been closed
//ok is false if there is no more values to receive and the channel is closed

// only the sender should close a channel
// sending on a closed channel will cause a panic

// closing channels is only necessary when the receiver must be told there are no more values coming
// such as 'range', it needs to know there are no more values to terminate the 'range' loop

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// we can range through a channel
	// doing this receives values from the channel until it is closed
	for i := range c {
		fmt.Println(i)
	}
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	// Closing the channel so the range loop can stop
	close(c)
}
