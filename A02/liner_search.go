package main

import (
	"fmt"
)

func main() {
	n, x := 0, 0
	fmt.Scan(&n, &x)

	for i := 0; i < n; i++ {
		target := 0
		fmt.Scan(&target)
		if target == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
