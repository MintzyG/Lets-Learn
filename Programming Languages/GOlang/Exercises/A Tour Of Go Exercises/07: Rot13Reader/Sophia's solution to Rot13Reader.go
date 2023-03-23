// I still have question about this one
package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

// Why passing 'b' by value is actually changing its value outside the function???
func (rot *rot13Reader) Read(b []byte) (int, error) {

	i, err := rot.r.Read(b)
	if err != nil {
		return 0, nil
	}

	for i := range b {

		if b[i] >= 65 && b[i] <= 91 {
			b[i] = 65 + (b[i]-52)%26
		} else if b[i] >= 97 && b[i] <= 123 {
			b[i] = 97 + (b[i]-84)%26
		}

	}

	return i, err

}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		return
	}
}
