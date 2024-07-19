package main

import "fmt"

func max(i, j int) int {
	if i < j {
		return j
	}
	return i
}

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	a := make([]int, n)
	b := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a[i], &b[i])
	}

	answer := 0
	for A := 1; A <= 100; A++ {
		for B := 1; B <= 100; B++ {
			counter := 0
			for i := 0; i < n; i++ {
				if A <= a[i] && a[i] <= A+k && B <= b[i] && b[i] <= B+k {
					counter++
				}
			}
			answer = max(answer, counter)
		}
	}

	fmt.Println(answer)
}
