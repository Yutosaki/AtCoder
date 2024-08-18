package main

import "fmt"

type jewelry struct {
	weight int
	value  int
}

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	n, w := 0, 0
	fmt.Scan(&n, &w)

	jewelrys := make([]jewelry, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&jewelrys[i].weight, &jewelrys[i].value)
	}

	dp := make([][]int, n+1) //i個の宝石で重さjを越えない様にする
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, w+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= w; j++ {
			if j < jewelrys[i].weight {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i-1][j-jewelrys[i].weight]+jewelrys[i].value)
			}
		}
	}

	fmt.Println(dp[n][w])
}
