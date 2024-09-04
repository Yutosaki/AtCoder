package main

import "fmt"

func max(a, b int) int {
	if a> b{
		return a
	}
	return b
}

func main() {
	s, t:="",""
	fmt.Scan(&s, &t)

	dp:=make([][]int, len(s)+1)
	for i:=0; i<len(s)+1; i++ {
		dp[i]=make([]int, len(t)+1)
	}

	for i:=1; i<=len(s); i++ {
		for j:=1; j<=len(t); j++ {
			if s[i-1]==t[j-1] {
				dp[i][j]=dp[i-1][j-1]+1
			} else {
				dp[i][j]=max(dp[i-1][j], dp[i][j-1])
			}
		}
	}

	fmt.Println(dp[len(s)][len(t)])
}
