package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	P := make([]int, n+2)
	A := make([]int, n+2)

	for i := 1; i <= n; i++ {
		fmt.Scan(&P[i], &A[i])
	}

	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}

	for len := n - 2; len >= 0; len-- {
		for l := 1; l <= n-len; l++ {
			r := l + len
			score1 := 0
			score2 := 0

			if l <= P[l-1] && P[l-1] <= r {
				score1 = A[l-1]
			}

			if l <= P[r+1] && P[r+1] <= r {
				score2 = A[r+1]
			}

			if l == 1 {
				dp[l][r] = dp[l][r+1] + score2
			} else if r == n {
				dp[l][r] = dp[l-1][r] + score1
			} else {
				dp[l][r] = max(dp[l-1][r]+score1, dp[l][r+1]+score2)
			}
		}
	}

	answer := 0
	for i := 1; i <= n; i++ {
		answer = max(answer, dp[i][i])
	}
	fmt.Println(answer)
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
