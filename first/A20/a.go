package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	S, T := "", ""
	fmt.Scan(&S, &T)

	dp := make([][]int, len(S)+1) //どのようなDPテーブルの定義にするかを考えてから始める。基本的には出力するものを格納していくべき
	for i := 0; i <= len(S); i++ {
		dp[i] = make([]int, len(T)+1)
	}

	for i := 1; i <= len(S); i++ {
		for j := 1; j <= len(T); j++ {
			if S[i-1] == T[j-1] {
				dp[i][j] = max(max(dp[i-1][j], dp[i][j-1]), dp[i-1][j-1]+1)
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}
	}
	fmt.Println(dp[len(S)][len(T)])
}
