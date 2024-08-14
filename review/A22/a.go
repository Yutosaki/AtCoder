package main

import (
	"fmt"
	"math"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	var N int
	fmt.Scan(&N)

	A := make([]int, N)
	B := make([]int, N)
	for i := 1; i < N; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i < N; i++ {
		fmt.Scan(&B[i])
	}

	dp := make([]int, N+1)
	
	dp[1] = 0
	for i := 2; i <= N; i++ {
		dp[i] = math.MinInt32
	}

	for i := 1; i < N; i++ {
		dp[A[i]] = max(dp[A[i]], dp[i]+100)
		dp[B[i]] = max(dp[B[i]], dp[i]+150)
	}

	fmt.Println(dp[N])
}