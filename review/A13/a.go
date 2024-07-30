package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	maxEnd := make([]int, n)
	counter := 0
	for i := 0; i < n-1; i++ {
		if i == 0 {
			maxEnd[i] = 0
		} else {
			maxEnd[i] = maxEnd[i-1]
		}

		for maxEnd[i] < n && A[maxEnd[i]]-A[i] <= k {
			maxEnd[i]++
		}
		counter += maxEnd[i] - i - 1
	}

	fmt.Println(counter)
}
