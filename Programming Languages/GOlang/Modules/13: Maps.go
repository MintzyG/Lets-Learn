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

	myCompanies := map[string]Companies{
		"Google": {
			"Google",
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

}
