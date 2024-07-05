package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n)
	B := make([]int, n)
	for i := 0; i < n-1; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < n-1; i++ {
		fmt.Scan(&B[i])
	}

	dp := make([]int, n+1)
	for i := 1; i <= n-1; i++ {
		dp[A[i-1]] = max(dp[A[i-1]], dp[i]+100)
		dp[B[i-1]] = max(dp[B[i-1]], dp[i]+150)
	}

	fmt.Println(dp[n])
}
