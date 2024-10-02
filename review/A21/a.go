package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	RemoveFrom := make([]int, n+2)
	Point := make([]int, n+2)

	for i := 1; i <= n; i++ {
		fmt.Scan(&RemoveFrom[i], &Point[i])
	}

	dp := make([][]int, n+2)
	for i := 0; i <= n+1; i++ {
		dp[i] = make([]int, n+2)
	}

	for length := n - 2; length >= 0; length-- {
		for left := 1; left <= n-length; left++ {
			right := left + length
			scoreLeft := 0
			scoreRight := 0

			if left <= RemoveFrom[left-1] && RemoveFrom[left-1] <= right {
				scoreLeft = Point[left-1]
			}

			if left <= RemoveFrom[right+1] && RemoveFrom[right+1] <= right {
				scoreRight = Point[right+1]
			}

			if left == 1 {
				dp[left][right] = dp[left][right+1] + scoreRight
			} else if right == n {
				dp[left][right] = dp[left-1][right] + scoreLeft
			} else {
				dp[left][right] = max(dp[left-1][right]+scoreLeft, dp[left][right+1]+scoreRight)
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
