package main

import "fmt"

func main() {
	n, target := 0, 0
	fmt.Scan(&n, &target)
	printers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&printers[i])
	}
	left, right := 1, 1000000000
	for left < right {
		mid := (left + right) / 2
		tmp := count(mid, printers, n, target)
		if tmp < target {
			left = mid + 1
		} else {
			right = mid
		}
	}

	fmt.Println(left)
}

func count(time int, printers []int, n int, target int) (counter int) {
	for i := 0; i < n; i++ {
		counter += time / printers[i]
		if counter > target {
			return
		}
	}
	return
}
