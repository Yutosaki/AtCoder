package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	mod := 1000000007
	answer := a
	index := 1

	for index*2 < b {
		answer = (answer * answer) % mod
		index *= 2
	}

	for index < b {
		answer = (answer * a) % mod
		index++
	}

	fmt.Println(answer)
}
