/*
Every GO program is made up of packages
The main package is where a program starts
Packages contain code to be executed, either in themselves or in other packages
For example you can import the math package in to main, to use its functionality
*/
package main

/*
Import groups imported packages in factored way
Imports gets the named packages into your current program, so you can use their functionality
Import can also be written:
import "fmt"
import "math/rand"
But it's not recommended to do so
*/
import (
	"fmt"
	"math"
	"math/rand"
)

/*
Exported names:
These are names that are exported from the imported functions
All exported names need to start with a capital letter, so:
math.Pi is the correct use and math.pi will give you an error

When importing a package, you can only refer to it's exported names
Any unexported names are not accessible from outside the imported package
*/
func main() {
	fmt.Println("My favorite number is: ", rand.Intn(10), " and ", math.Pi)
}
