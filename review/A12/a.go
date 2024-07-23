package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	left := 0
	right := 100000000
	for left < right {
		mid := (left + right) / 2
		if printNum(mid, A) < k {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}

func printNum(time int, array []int) int {
	counter := 0
	for i := 0; i < len(array); i++ {
		counter += time / array[i]
	}
	return counter
}
