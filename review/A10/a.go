package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	n := 0

	fmt.Scan(&n)
	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	leftMax := make([]int, n+1)
	for i := 1; i <= n; i++ {
		leftMax[i] = max(leftMax[i-1], A[i])
	}
	rightMax := make([]int, n+1)
	rightMax[n] = A[n]
	for i := n - 1; i >= 1; i-- {
		rightMax[i] = max(rightMax[i+1], A[i])
	}
	d := 0
	fmt.Scan(&d)
	ans := make([]int, d)
	l, r := 0, 0
	for i := 0; i < d; i++ {
		fmt.Scan(&l, &r)
		ans[i] = max(leftMax[l-1], rightMax[r+1])
	}

	for i := 0; i < d; i++ {
		fmt.Println(ans[i])
	}
}
