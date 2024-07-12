package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	xorSum := A[0]
	for i := 1; i < n; i++ {
		xorSum ^= A[i]
	}

	if xorSum != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
