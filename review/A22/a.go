package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n)
	B := make([]int, n)

	for i := 1; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 1; i < n; i++ {
		fmt.Scan(&B[i])
	}

	dp := make([]int, n+1)
	for i:=0;i<n;i++{
		dp[i]=-2147483648
	}

	dp[1]=0
	for i := 1; i < n; i++ {
		dp[A[i]] = max(dp[A[i]], dp[i]+100)
		dp[B[i]] = max(dp[B[i]], dp[i]+150)
	}

	fmt.Println(dp[n])
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
