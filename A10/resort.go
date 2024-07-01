package main

import (
	"fmt"
	"sort"
)

type number struct {
	day      int
	capacity int
}

func main() {
	n := 0
	fmt.Scan(&n)
	numbers := make([]number, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&numbers[i].capacity)
		numbers[i].day = i
	}
	sort.Slice(numbers, func(i, j int) bool { return numbers[i].capacity > numbers[j].capacity })

	d := 0
	fmt.Scan(&d)
	answers := make([]int, d)
	for i := 0; i < d; i++ {
		l, r := 0, 0
		fmt.Scan(&l, &r)
		for j := 0; j <= n; j++ {
			if numbers[j].day < l || numbers[j].day > r {
				answers[i] = numbers[j].capacity
				break
			}
		}
	}

	for i := 0; i < d; i++ {
		fmt.Println(answers[i])
	}
}
