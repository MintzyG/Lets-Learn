package main

import "fmt"

// Let's say we have an int named number with value 4, and we want to print it
// We use fmt.Println(number), pretty easy right, now lets create our own type and try to print it

type Customer struct {
	Name     string
	Age      int
	Customer bool
}

func main() {

	Customer1 := Customer{"John", 32, true}

	// If we try that same line of code, but with our type we get

	fmt.Println(Customer1)

	// We get this '{John 32 true}', which is okay, but we can make it better
	// By using an interface called Stringer we can tell our code how stuff should be printed
	// For example we could make our type Person print out in a better way than what we got

	// First we set up the Stringer just like we set up any other interface
	// But this time we need to call the method String which is already defined by the stringer interface
	// This method returns a string, so it can be printed
	/*
		func (p Person) String() string {
			var text string
			if p.Customer {
				text = "is"
			} else {
				text = "is not"
			}

			return fmt.Sprintf("%v is %v year old and %v a customer", p.Name, p.Age, text)
		}
	*/
	// This stringer tells our code how to print our type Person
	// Let's test it out

	fmt.Println(Customer1)

	// Nice it works, and don't mind both outputs being equal, since this change takes effect everywhere in the code
	// So even our first print is going to print it right now, it only printed that weird way, before we implemented the stringer
}

func (c Customer) String() string {
	var text string
	if c.Customer {
		text = "is"
	} else {
		text = "is not"
	}

	return fmt.Sprintf("%v is %v year old and %v a customer", c.Name, c.Age, text)
}
