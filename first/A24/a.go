package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	long := make([]int, n+1)
	len := 0

	for i := 1; i <= n; i++ {
		pos := sort.SearchInts(long[1:len+1], A[i]) + 1

		long[pos] = A[i]
		if pos > len {
			len++
		}
	}
	fmt.Println(len)
}
