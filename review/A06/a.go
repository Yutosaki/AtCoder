package main

import "fmt"

func main() {
	n, q := 0, 0
	fmt.Scan(&n, &q)

	tmp := 0
	subsetSum := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&tmp)
		subsetSum[i] = tmp + subsetSum[i-1]
	}

	ans := make([]int, q)
	right, left := 0, 0
	for i := 0; i < q; i++ {
		fmt.Scan(&left, &right)
		ans[i] = subsetSum[right] - subsetSum[left-1]
	}

	for i := 0; i < q; i++ {
		fmt.Println(ans[i])
	}
}
