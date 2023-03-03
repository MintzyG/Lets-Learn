package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	var Bubble [20]int
	var Swap int
	var Sorted bool

	rand.Seed(time.Now().Unix())

	for i := range Bubble {
		Bubble[i] = rand.Intn(100) + 1
	}

	fmt.Println("Unsorted array:", Bubble)

	for {

		for i := 0; i < len(Bubble)-1; i++ {

			if Bubble[i] > Bubble[i+1] {
				Sorted = false
				break
			} else {
				Sorted = true
			}

		}

		if !Sorted {
			for i := 0; i < len(Bubble)-1; i++ {
				if Bubble[i] > Bubble[i+1] {
					Swap = Bubble[i]
					Bubble[i] = Bubble[i+1]
					Bubble[i+1] = Swap
				}
			}
		} else {
			break
		}

	}

	fmt.Println(Bubble)

}
