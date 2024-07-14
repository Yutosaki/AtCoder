package main

import "fmt"

func main() {
	n, diff := 0, 0
	fmt.Scan(&n, &diff)
	array := make([]int, n)

	counter := 0
	for i := 0; i < n; i++ {
		fmt.Scan(&array[i])
	}

	j := 0
	for i := 1; i < n; i++ {
		for j < i && array[i]-array[j] > diff {
			j++
		}
		counter += i - j
	}
	fmt.Println(counter)
}
