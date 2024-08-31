package main

import (
	"fmt"
	"sort"
)

func main() {
	n, target := 0, 0
	fmt.Scan(&n, &target)

	A := make([]int, n)
	B := make([]int, n)
	C := make([]int, n)
	D := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&B[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&C[i])
	}
	for i := 0; i < n; i++ {
		fmt.Scan(&D[i])
	}

	P := make([]int, n*n)
	Q := make([]int, n*n)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			P[i*n+j] = A[i] + B[j]
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Q[i*n+j] = C[i] + D[j]
		}
	}

	sort.Slice(Q, func(i, j int) bool { return Q[i] < Q[j] })

	for i := 0; i < n*n; i++ {
		if findPlace(target-P[i], Q) {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}

func findPlace(target int, array []int) bool {
	left := 0
	right := len(array) - 1
	for left <= right {
		mid := (left + right) / 2
		if array[mid] == target {
			return true
		} else if array[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
