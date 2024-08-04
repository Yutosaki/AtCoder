package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	n, W := 0, 0
	fmt.Scan(&n, &W)
	w := make([]int, n+1)
	v := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&w[i], &v[i])
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if j < w[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-w[i]]+v[i])
			}
		}
	}

	fmt.Println(dp[n][W])
}
