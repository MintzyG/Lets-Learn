package main

import "fmt"

func main() {

	// A struct is a collection of fields
	// This collection of fields describe the struct
	// For example I have a Person struct, a collection of fields for that struct could be: string Name, int Age
	type Vertex struct {
		X int
		Y int
	}

	// The fields of a struct are accessed by using the dot notation
	// But for that the struct has to be initialized like a variable
	// Although you don't need to assign variables to its fields to access them

	// Accessing a struct field without having assigned variables on it
	// Any unassigned field will be zero valued
	v0 := Vertex{}
	v0.X = 5
	fmt.Println(v0)

	// You can create a pointer to a struct
	p := &v0

	// You could access the struct fields through the pointer using (*p).X, but that is impractical
	// So if you have a struct pointer you are allowed to omit the *
	p.Y = 8
	fmt.Println(v0)

	// Struct Literals
	// A struct literal describes a newly allocated struct by listing the values of its fields
	// You can also declare only certain parts of a struct by specifying its name EX: Vertex{X: 3} only changes the x value
	// Using & prefix returns a pointer to the struct value
	v1 := Vertex{1, 2}  // X:1 and Y:2
	v2 := Vertex{X: 1}  // X:1 and Y:0
	v3 := Vertex{Y: 2}  // X:0 and Y:2
	v4 := Vertex{}      // X and Y are 0
	p1 := &Vertex{1, 2} // Has type *Vertex

	fmt.Println(v1, v2, v3, v4, p1, *p1)

	// Embeds
	// read the embeds section below
	embeds()

}

// Embedding structs

// Base struct and their Stringer method
type base struct {
	num int
}

func (b base) describe() string {
	return fmt.Sprintf("base with num = %v", b.num)
}

// A container embeds a base
// Embedding looks like a field without a name

type container struct {
	base
	str string
}

func embeds() {

	// When creating structs with literals we have to initialize the embed explicitly
	// Here the embedded type serves as the field name
	co := container{
		base: base{
			num: 1,
		},
		str: "Container 1",
	}

	// With this embedded type we can access the base's fields directly on 'co'

	fmt.Printf("co={num: %v, str: %v}\n", co.num, co.str)

	// Or we can spell out the full path using the embedded type name

	fmt.Println("Also num:", co.base.num)

	// Since container embeds 'base', its methods also become methods of container
	// Here we invoke a method that was embedded from base directly onto 'co'

	fmt.Println("Describe:", co.describe())

	// Embedding structs with methods may be used to bestow interface implementations in to other structs
	// Here we see that container implements the 'describer' interface because it embeds 'base'

	type describer interface {
		describe() string
	}

	var d describer = co
	fmt.Println("Describer:", d.describe())
}
