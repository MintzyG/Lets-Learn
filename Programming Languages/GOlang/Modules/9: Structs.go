package main

import "fmt"

func main() {

	// A struct is a collection of fields
	// These collection of fields describe the struct
	// For example I have a Person struct, a collection of fields for that struct could be: string Name, int Age
	type Vertex struct {
		X int
		Y int
	}

	// The fields of a struct are accessed by using the dot notation
	// But for that the struct has to be initialized like a variable
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	// You can create a pointer to a struct
	p := &v

	// You could access the struct fields through the pointer using (*p).X, but that is impractical
	// So if you have a struct pointer you are allowed to omit the *
	p.Y = 8
	fmt.Println(v)

	// Struct Literals
	// A struct literal describes a newly allocated struct by listing the values of its fields
	// You can also just list a subset of fields by using 'field': syntax (and now the order is irrelevant)
	// Using & prefix returns a pointer to the struct value
	v1 := Vertex{1, 2}  // X:1 and Y:2
	v2 := Vertex{X: 1}  // X:1 and Y:0
	v3 := Vertex{Y: 2}  // X:0 and Y:2
	v4 := Vertex{}      // X and Y are 0
	p1 := &Vertex{1, 2} // Has type *Vertex

	fmt.Println(v1, v2, v3, v4, p1, *p1)
}
