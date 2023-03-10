package main

import (
	"fmt"
	"math"
)

// Interfaces are named collections of method signatures
// Meaning they are a type that list all methods that have to be implemented
// For any given type for it to  be considered a valid interface type
// Here's a basic interface for geometric shapes

type geometry interface {
	area() float64
	perim() float64
}

// This interface describes the methods that must be implemented on a type
// for it to be considered a 'geometry'

// Here are two types we will implement our interfaces on

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

// Let's implement the 'geometry' interface on them
// For that we need to implement all methods described in the interface on our types

// Let's do that for rect

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

// Now rect implements 'geometry' since it has all the methods it describes
// Let's do circle now

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// Circle now implements 'geometry' as well
// Notice how both types implement 'geometry' but their methods are completely different
// This shows that it doesn't matter how each method is made just that all methods need to be implemented to satisfy the interface

// Now that we have types with our interface implemented we can call the methods from the interface through these types
// But we can take advantage of the interface type and create a generic function that takes it
// So that we can pass any type that implements 'geometry' to it
// Here's a generic function that works with any 'geometry'

func measure(g geometry) {
	fmt.Println("The perimeter is", g.perim())
	fmt.Println("The area is", g.area())
}

// Let's test it out

func main() {

	r := rect{width: 3, height: 4}
	c := circle{radius: 8}

	measure(r)
	measure(c)

}

//TODO: reimplement A-Tour-of_GO's interface explanations including
// type assertions and type switches
