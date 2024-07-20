package main

import "fmt"

func printCount(currenttime int, A []int) int {
	ans := 0
	for i := 0; i < len(A); i++ {
		ans += currenttime / A[i]
	}
	return ans
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	left := 1
	right := 100000000
	for left < right {
		mid := (left + right) / 2
		if printCount(mid, A) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}
