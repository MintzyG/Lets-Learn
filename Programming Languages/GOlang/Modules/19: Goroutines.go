package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

//FIXME: For some reason goroutines are broken, find out later and fix it/explain it

func main() {

	f("Function call")

	go f("Goroutine")

	time.Sleep(time.Second)

}
