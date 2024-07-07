package main

import "fmt"

func main() {
	n, s := 0, 0
	fmt.Scan(&n, &s)
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	dp := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = make([]bool, s+1)
	}

	dp[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if j < A[i] {
				dp[i][j] = dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j] || dp[i-1][j-A[i]]
			}
		}
	}

	if dp[n][s] == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
