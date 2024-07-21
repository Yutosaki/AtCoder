package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	counter := 0
	for i := 0; i < n; i++ {
		left := i + 1
		right := n

		// 二分探索でA[i] + K以下の最大の位置を見つける
		for left < right {
			mid := (left + right) / 2
			if A[mid] <= A[i] + k {
				left = mid + 1
			} else {
				right = mid
			}
		}

		// 見つけた位置leftはA[i] + K以下の最大の位置 + 1
		counter += left - (i + 1)
	}

	fmt.Println(counter)
}
