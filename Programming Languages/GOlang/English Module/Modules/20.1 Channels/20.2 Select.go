package main

import "fmt"

// The select statement lets the goroutine wait on multiple channels
// The 'select' keyword blocks until one of its cases can run, then it executes that case
// If multiple are ready it chooses at random

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	Fibonacci(c, quit)
}

func Fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {

		// Select waits for a channel to be ready to execute
		// When the channel can execute it does so
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}
