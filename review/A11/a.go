package main

import "fmt"

func main() {
	n, x := 0, 0
	fmt.Scan(&n, &x)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	left := 0
	right := n - 1
	for left < right {
		mid := (left + right) / 2
		if A[mid] < x {
			left = mid + 1
		} else if A[mid] > x {
			right = mid - 1
		} else {
			left = mid
			break
		}
	}

	fmt.Println(left + 1)
}
