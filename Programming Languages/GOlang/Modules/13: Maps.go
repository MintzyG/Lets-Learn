package main

import "fmt"

func main() {

	// A 'map' maps keys to values
	// The zero value of a map is nil, A nil map has no keys, and they can't be added
	// The make() function returns a map of the given type initialized

	type Vertex struct {
		X int
		Y int
	}

	type Companies struct {
		Name    string
		Founded int // Just the year
	}

	// declaring a map
	// A map is declared by typing map, then inside [] its key type, then what it maps to
	// map[key Type]MapsTo
	myMap := make(map[string]Vertex)

	myMap["First"] = Vertex{
		1,
		2,
	}
	myMap["Second"] = Vertex{
		2,
		4,
	}

	fmt.Println(myMap["Second"])

	// Ranges through all the keys in the map returning the value on that key
	for _, v := range myMap {

		fmt.Println(v.X, v.Y)

	}

	// Map literals:
	// Just like arrays, structs and slices maps can be declared through literals

	// If the top level type is just a type name you can omit it from the elements of the literal
	// Meaning I didn't have to type 'Companies' alongside the keys because the top level type here is a type name
	myCompanies := map[string]Companies{
		"Google": {
			"Alphabet",
			1989,
		},
		"Amazon": {
			"Amazon",
			1995,
		},
	}

	for _, v := range myCompanies {

		fmt.Println(v.Name, "was founded in", v.Founded)

	}

	// Mutating maps:

	// You can insert or update an element in the map by using:
	// m[Key] = element
	// For example:

	// Inserting an element in the map
	myCompanies["Apple"] = Companies{
		"Apple",
		1975,
	}

	// Updating an element in the map e
	myCompanies["Google"] = Companies{
		"Google",
		1989,
	}

	fmt.Println("-------------------")

	for _, v := range myCompanies {

		fmt.Println(v.Name, "was founded in", v.Founded)

	}

	// Retrieving an element from a map
	elem := myCompanies["Apple"]
	fmt.Println(elem)

	// Delete an element from a map
	delete(myCompanies, "Google")

	// Test if an element is present with a two value assignment
	// If the key is in the map 'ok' is true,if not 'ok' is false
	// If the key is not in the map element is its zero type
	elem2, ok := myCompanies["Google"]
	if ok {
		fmt.Println(elem2)
	} else {
		fmt.Println("Element was removed")
		fmt.Println("Element's zero type:", elem2)
	}

	// Note if elem or ok have not been declared you can use the short declaration just like in the example above

}
