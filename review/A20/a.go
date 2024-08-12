package main

import (
	"bufio"
	"fmt"
	"os"
)

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	S := scanner.Text()
	scanner.Scan()
	T := scanner.Text()

	n := len(S)
	m := len(T)

	dp := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[n][m])
}
