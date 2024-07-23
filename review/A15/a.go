package main

import (
	"fmt"
	"sort"
)

type numbers struct {
	base        int
	value       int
	compression int
}

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]numbers, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i].value)
		A[i].base = i
	}

	sort.Slice(A, func(i, j int) bool { return A[i].value < A[j].value })

	index := 1
	A[0].compression = index
	for i := 1; i < n; i++ {
		if A[i-1].value != A[i].value {
			index++
		}
		A[i].compression = index
	}

	sort.Slice(A, func(i, j int) bool { return A[i].base < A[j].base })

	for i := 0; i < n; i++ {
		fmt.Printf("%d ", A[i].compression)
	}
	fmt.Println()
}
