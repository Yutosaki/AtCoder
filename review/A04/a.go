package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	binary := ""
	for i := 0; i < 10; i++ {
		binary = fmt.Sprintf("%d", n%2) + binary
		n /= 2
	}

	fmt.Println(binary)
}
