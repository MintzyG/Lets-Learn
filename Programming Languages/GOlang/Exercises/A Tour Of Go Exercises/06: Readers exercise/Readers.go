// Code only runs is 'A tour of go'
package main

import "golang.org/x/tour/reader"

type MyReader struct{}

func (m MyReader) Read(b []byte) (int, error) {

	b[0] = 65

	return 1, nil
}

func main() {
	reader.Validate(MyReader{})
}
