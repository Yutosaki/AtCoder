package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	A := make([]int, n+1)
	B := make([]int, n+1)

	for i := 2; i <= n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Scan(&B[i])
	}

	dp := make([]int, n+1)
	dp[2] = A[2]
	for i := 3; i <= n; i++ {
		if dp[i-1]+A[i] < dp[i-2]+B[i] {
			dp[i] = dp[i-1] + A[i]
		} else {
			dp[i] = dp[i-2] + B[i]
		}
	}

	revRoot := make([]int, 0)
	revRoot = append(revRoot, n)

	for i := n; i >= 1; {
		if dp[i] == dp[i-1]+A[i] {
			revRoot = append(revRoot, i-1)
			i--
		} else {
			revRoot = append(revRoot, i-2)
			i -= 2
		}
	}
	fmt.Println(len(revRoot) - 1)

	for i := len(revRoot) - 2; i >= 0; i-- {
		fmt.Print(revRoot[i])
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
