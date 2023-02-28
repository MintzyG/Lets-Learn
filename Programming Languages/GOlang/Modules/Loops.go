package main

import "fmt"

func main() {

	// Go has only one type of loop, the 'for' loop
	// A 'for' loop has three components separated by semicolons
	// Init statement ; Condition expression ; Post statement
	// Init will often be a short declaration of variables that are visible only inside the 'for' code scope
	// The loop will stop when the condition expression is equals to false
	// There are no parenthesis surrounding the three components of the 'for' loop and the {} are always required
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// Om top of that the init and post statements are optional

	sum = 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)

	// At this point you can remove the condition expression entirely and have a forever loop

	sum = 1
	for {
		sum += sum
		if sum > 3000 {
			break
		}
	}
	fmt.Println(sum)

}
