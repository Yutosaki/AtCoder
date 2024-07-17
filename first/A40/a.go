package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	sort.Ints(A)

	changes := 0
	current := A[0]
	for i := 1; i < n; i++ {
		if current != A[i] {
			changes++
		}
	}

	nums := make([]int, changes)
	index := 0
	current = A[0]
	for i := 1; i < n; i++ {
		if current == A[i] {
			nums[index]++ //一個少なくカウント
		} else {
			current = A[i]
			index++
		}
	}

	ans := 0
	for i := 0; i < changes; i++ {
		if nums[i]+1 >= 3 {
			ans += nC3(nums[i] + 1)
		}
	}

	fmt.Println(ans)
}

func nC3(n int) int {
	return (n * (n - 1) * (n - 2)) / 6
}
