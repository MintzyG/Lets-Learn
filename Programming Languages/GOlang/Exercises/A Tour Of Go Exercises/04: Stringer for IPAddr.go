// Very bad solution, but a future proof one
package main

import "fmt"
import "strconv"

type IPAddr [4]byte

func (I IPAddr) String() string {
	var IP string
	var str string
	for i, v := range I {
		str = strconv.Itoa(int(v))
		if i != 3 {
			IP += str + "."
		} else {
			IP += str
		}
	}

	return fmt.Sprintf(IP)
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
