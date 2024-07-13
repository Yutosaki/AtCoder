package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func min(i, j int) int {
	if i > j {
		return j
	}
	return i
}

func main() {
	n := 0
	fmt.Scan(&n)
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	dp := make([][]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		dp[n][i] = A[i]
	}

	for i := n - 1; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			if i%2 == 1 {
				dp[i][j] = max(dp[i+1][j], dp[i+1][j+1])
			} else {
				dp[i][j] = min(dp[i+1][j], dp[i+1][j+1])
			}
		}
	}
	fmt.Println(dp[1][1])
}
