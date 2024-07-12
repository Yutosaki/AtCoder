package main

import "fmt"

func main() {
	n, X, Y := 0, 0, 0
	fmt.Scan(&n, &X, &Y)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	grundy := make([]int, 100001)

	for i := 0; i <= 100000; i++ {
		transit := [3]bool{false, false, false}
		if i >= X {
			transit[grundy[i-X]] = true
		}
		if i >= Y {
			transit[grundy[i-Y]] = true
		}
		if !transit[0] {
			grundy[i] = 0
		} else if !transit[1] {
			grundy[i] = 1
		} else {
			grundy[i] = 2
		}
	}

	XOR_Sum := 0
	for i := 0; i < n; i++ {
		XOR_Sum ^= grundy[A[i]]
	}
	if XOR_Sum != 0 {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
