package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)

	for a != 0 {
		a, b = bigFirst(a, b)
		a = a % b
	}
	fmt.Println(b)
}

func bigFirst(a, b int) (int, int) {
	if a < b {
		a, b = b, a
	}
	return a, b
}
