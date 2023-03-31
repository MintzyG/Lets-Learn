package main

import (
	"fmt"
)

type ErrNegNum float64

func (e ErrNegNum) Error() string {
	return fmt.Sprintf("Cannot square root the negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {

	if x < 0 {
		return 0, ErrNegNum(x)
	} else if x == 0 {
		return 0, nil
	}

	z := x - x/2

	for {
		z -= (z*z - x) / (2 * z)

		fmt.Println(z*z - x)

		if z*z-x < 0.1 {
			return z, nil
		}
	}
}

func main() {
	var enter float64 = 1234
	for {
		v, err := Sqrt(enter)
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Printf("%.2f is the sqrt of %v", v, enter)
		break
	}
}
