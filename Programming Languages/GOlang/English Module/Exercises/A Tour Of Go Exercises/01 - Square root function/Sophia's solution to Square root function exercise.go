package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := x - x/2
	for {
		z -= (z*z - x) / (2 * z)

		fmt.Println(z*z - x)

		if z*z-x < 0.1 {
			return z
		}
	}
}

func main() {
	fmt.Println(Sqrt(100))
}
