package main

import (
	"fmt"
	"math"
)

func min(i, j int) int {
	if i < j {
		return i
	}
	return j
}

func main() {
	n, m := 0, 0
	fmt.Scan(&n, &m)

	A := make([][]int, m+1)
	for i := 1; i <= m; i++ {
		A[i] = make([]int, n+1)
		for j := 1; j <= n; j++ {
			fmt.Scan(&A[i][j])
		}
	}

	dp := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]int, 1<<n)
		for j := 1; j < (1 << n); j++ {
			dp[i][j] = math.MaxInt32
		}
	}

	dp[0][0] = 0

	for i := 1; i <= m; i++ {
		for j := 0; j < (1 << n); j++ {
			v := j
			for k := 1; k <= n; k++ {
				if A[i][k] == 1 {
					v |= (1 << (k - 1))
				}
			}
			dp[i][j] = min(dp[i][j], dp[i-1][j])
			dp[i][v] = min(dp[i][v], dp[i-1][j]+1)
		}
	}

	if dp[m][(1<<n)-1] == math.MaxInt32 {
		fmt.Println("-1")
	} else {
		fmt.Println(dp[m][(1<<n)-1])
	}
}
