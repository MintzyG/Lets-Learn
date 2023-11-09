package main

/*

This code does not run in an IDE

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	picture := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		picture[i] = make([]uint8, dx)
		for j := 0; j < dx; j++ {
			picture[i][j] = uint8((j ^ i + (i ^ j)) * (j ^ 2 + i ^ 2) / 50)
		}
	}

	return picture
}

func main() {
	pic.Show(Pic)
}
*/
