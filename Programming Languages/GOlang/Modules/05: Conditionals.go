package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {

	// The 'if' is like a 'for', you don't need the parenthesis but the {} are required

	sum := 1
	for {
		sum += sum
		if sum > 32 {
			fmt.Println("Sum has overflown")
			break
		}
		fmt.Println(sum)
	}

	// The 'if' can also have a short statement before its condition
	// Variables declared inside the 'if' short statement are also available in the else blocks

	sum = 1
	decrease := 10
	for {
		sum += sum
		decrease -= 1
		if v := sum * decrease; sum > v {
			fmt.Println("Sum has overflown")
			break
		} else {
			fmt.Println("Value: ", v)
		}
		fmt.Println("Sum : ", sum)
	}

	// Switches:
	// A shorter way to write 'if - else' blocks
	// It runs the first case whose value is equal to the condition

	// Go only runs the selected case, so no break is needed, since it won't run all following cases
	// Values used for evaluation don't need to be constants nor integers, can be strings, floats, etc

	// Switches in go evaluate cases from top to bottom, so if a case succeeds, cases below won't be executed

	fmt.Println("Running on: ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}

	// Switches can have no condition statement
	// Making it with a good structure can create 'if - then - else' like block
	// Example: if -> (time < 12) then -> {"good"} else -> if (time < 17) ... etc
	// Switches can use commas to separate multiple expressions in a single case

	t := time.Now()
	switch {
	default:
		fmt.Println("Good evening.")
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17, t.Hour() < 18:
		fmt.Println("Good afternoon.")
	}

}
