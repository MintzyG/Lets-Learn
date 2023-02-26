package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

/*
GO's basic types are:

Bool: Equivalent to true or false, 0 or 1

String: Stores text

int int8 int16 int32 int64: Stores numbers, the numbers in the int name signify how many bits long it can hold
int is a signed type, meaning it can hold negative values using its leftmost bit as the sign
Example: int8 would be 8 bits long so -> 01111111 is the biggest positive number it can store
int without a number is default for int32

uint uint8 uint16 uint32 uint64 uintptr: Same as int, but the u is for unsigned, meaning it doesn't support
	negative numbers effectively doubling the biggest number it can store
Example: uint8 -> 11111111 is the biggest number it can store
uint without the number is default for uint32

byte: alias for uint8
basically it functions the same way as uint8

rune: alias for int32
	Represents a unicode point, meaning, a unicode character

float32, float64: similar to int but it supports floating points meaning it supports decimal points like:
	2.58732 or 32456.342
floating point because the decimal point can be anywhere in the number like seen in the example above

complex64 complex128: used to store complex numbers
*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)

	// Variables declared without an explicit initial value are given their zero value
	// 0 for numeric types, false for boolean, "" empty string for strings
	var i int
	var f float64
	var b bool
	var s string
	fmt.Println(i, f, b, s)

	// You can convert types by using T(v), this converts the variable v to the type T
	// Conversions need to be explicit
	var number int = 32
	var float float64 = math.Sqrt(float64(number))
	var u uint = uint(float)
	// if done inside a function you could use the short assignment
	// Example float := math.Sqrt(float64(number))

	fmt.Println(number, float, u)

}
