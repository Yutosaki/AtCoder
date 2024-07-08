package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scan(&a, &b)
	a, b = bigFirst(a, b)

	for b != 0 {
		a, b = bigFirst(a, b)
		a -= b
	}
	fmt.Println(a)
}

func bigFirst(a, b int) (int, int) {
	if a < b {
		tmp := b
		b = a
		a = tmp
	}
	return a, b
}
