package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	counter := 0
	end := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if A[i] == 1 {
			end[i] = 1
		} else {
			end[i] = end[i-1]
		}

		for end[i] <= n && A[end[i]]-A[i] <= k {
			end[i]++
		}
		counter += end[i] - i - 1
	}

	fmt.Println(counter)
}
