package main

import (
	"fmt"
)

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	mod := 1000000007
	answer := 1
	power := a

	for i := 0; i < 30; i++ {
		wari := (1 << i)
		if (b/wari)%2 == 1 {
			answer = (answer * power) % mod
		}
		power = (power * power) % mod
	}

	fmt.Println(answer)
}
