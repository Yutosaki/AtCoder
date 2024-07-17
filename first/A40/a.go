package main

import (
	"fmt"
)

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	count := make([]int, 101)
	for i := 0; i < n; i++ {
		count[A[i]]++
	}

	ans := 0
	for i := 1; i <= 100; i++ {
		ans += nC3(count[i])
	}

	fmt.Println(ans)
}

func nC3(n int) int {
	return (n * (n - 1) * (n - 2)) / 6
}
