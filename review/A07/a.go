package main

import "fmt"

func main() {
	d, n := 0, 0
	fmt.Scan(&d, &n)
	array := make([]int, d+1)
	l, r := 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r)
		array[l-1]++
		array[r]--
	}

	for i := 1; i <= d; i++ {
		array[i] += array[i-1]
	}

	for i := 0; i < d; i++ {
		fmt.Println(array[i])
	}
}
