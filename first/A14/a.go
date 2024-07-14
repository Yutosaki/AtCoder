package main

import (
	"fmt"
	"sort"
)

func main() {
	n, sum := 0, 0
	fmt.Scan(&n, &sum)
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
	AB := make([]int, n*n)
	CD := make([]int, n*n)

	index := 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			AB[index] = A[i] + B[j]
			index++
		}
	}
	index = 0
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			CD[index] = C[i] + D[j]
			index++
		}
	}

	sort.Ints(CD)

	for _, v := range AB {
		left := 0
		right := n*n - 1

		for left < right {
			mid := (left + right) / 2
			if v+CD[mid] < sum {
				left = mid + 1
			} else if v+CD[mid] > sum {
				right = mid - 1
			} else {
				fmt.Println("Yes")
				return
			}
		}
	}
	fmt.Println("No")
}
