package main

import (
	"fmt"
	"sort"
)

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

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
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			P = append(P, A[i]+B[j])
		}
	}

	Q := make([]int, 0, n*n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			Q = append(Q, C[i]+D[j])
		}
	}

	sort.Ints(Q)

	found := false
	for _, p := range P {
		target := k - p
		if binarySearch(Q, target) {
			found = true
			break
		}
	}

	if found {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}

func binarySearch(array []int, x int) bool {
	left := 0
	right := len(array) - 1

	for left <= right {
		mid := left + (right-left)/2
		if array[mid] == x {
			return true
		} else if array[mid] < x {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return false
}
