package main

import (
	"fmt"
)

func main() {
	n, target := 0, 0
	fmt.Scan(&n, &target)
	array := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	left := 0
	right := len(array) - 1
	for left < right {
		mid := (left + right) / 2
		if array[mid] < target {
			left = mid + 1
		} else if array[mid] > target {
			right = mid - 1
		} else {
			left = mid
			break
		}
	}
	fmt.Println(left + 1)
}
