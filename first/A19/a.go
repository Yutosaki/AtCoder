package main

import "fmt"

type jewelry struct {
	weight int
	value  int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n, W := 0, 0
	fmt.Scan(&n, &W)
	jewelrys := make([]jewelry, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&jewelrys[i].weight, &jewelrys[i].value)
	}

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= W; j++ {
			if j < jewelrys[i].weight {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-jewelrys[i].weight]+jewelrys[i].value)
			}
		}
	}

	fmt.Println(dp[n][W])
}
