package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	counter := 0
	for i := 0; i < n; i++ {
		left := i + 1
		right := n
		for left < right {
			mid := (left + right) / 2
			if A[mid]-A[i] <= k {
				left = mid + 1
			} else {
				right = mid
			}
		}
		counter += left - i - 1
	}

	fmt.Println(counter)
}
