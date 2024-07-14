package main

import "fmt"

type block struct {
	under_target int
	score        int
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	n := 0
	fmt.Scan(&n)

	blocks := make([]block, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&blocks[i].under_target, &blocks[i].score)
	}

	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}

	for length := n - 1; length >= 0; length-- {
		for l := 1; l <= n-length; l++ {
			r := l + length

			score1 := 0
			if l <= blocks[l-1].under_target && blocks[l-1].under_target <= r {
				score1 = blocks[l-1].score
			}

			score2 := 0
			if l <= blocks[r].under_target && blocks[r].under_target <= r {
				score2 = blocks[r].score
			}

			if l == 1 {
				dp[l][r] = dp[l][r+1] + score2
			} else if r == n {
				dp[l][r] = dp[l-1][r] + score1
			} else {
				dp[l][r] = max(dp[l-1][r], dp[l][r+1])
			}
		}
	}

	max_score := 0
	for i := 1; i <= n; i++ {
		max_score = max(max_score, dp[i][i])
	}
	fmt.Println(max_score)
}
