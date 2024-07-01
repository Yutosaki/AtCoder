package main

import (
	"fmt"
)

func max(a,b int)int{
	if a<b{
		return b
	}
	return a
}

func main() {
	n := 0
	fmt.Scan(&n)
	array := make([]int, n+1)
	front := make([]int, n+1)
	back := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&array[i])
	}

	front[1] = array[1]
	for i := 2; i <= n; i++ {
		front[i] = max(array[i], front[i-1])
	}

	back[n] = array[n]
	for i := n - 1; i >= 1; i-- {
		back[i] = max(array[i], back[i+1])
	}

	d := 0
	fmt.Scan(&d)
	for i := 0; i < d; i++ {
		l, r := 0, 0
		fmt.Scan(&l, &r)
		fmt.Println(max(front[l-1],back[r+1]))
	}
}
