package main

import (
	"fmt"
	"time"
)

// The select statement lets the goroutine wait on multiple channel operations
// The 'select' keyword blocks until one of its cases can run, then it executes that case
// If multiple are ready it chooses at random

func main() {

	// Let's select across this two channels
	c1 := make(chan int)
	c2 := make(chan int)

	// Each channel will receive a value after some amount of time to simulate blocking
	go func() {
		time.Sleep(1 * time.Second)
		c1 <- 1
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- 2
	}()

	// We'll use 'select' to wait for these values simultaneously, printing each one as it arrives
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("Received:", msg1)
		case msg2 := <-c2:
			fmt.Println("Received:", msg2)
		}
	}
}
