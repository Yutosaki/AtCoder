package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	P := make([]int, n+1)
	Q := make([]int, n+1)
	R := make([]int, n+1)

	for i := 1; i <= n; i++ {
		fmt.Scan(&P[i], &Q[i], &R[i])
	}

	X := make([]int, 21)
	for i := 1; i <= n; i++ {
		countA := 0
		countB := 0
		X[P[i]]++
		X[Q[i]]++
		X[R[i]]++
		for j := 1; j <= 20; j++ {
			if X[j] == 0 {
				countA++
			}
		}
		X[P[i]] -= 2
		X[Q[i]] -= 2
		X[R[i]] -= 2
		for j := 1; j <= 20; j++ {
			if X[j] == 0 {
				countB++
			}
		}
		X[P[i]]++
		X[Q[i]]++
		X[R[i]]++

		if countA > countB {
			fmt.Println("A")
			X[P[i]]++
			X[Q[i]]++
			X[R[i]]++
		} else {
			fmt.Println("B")
			X[P[i]]--
			X[Q[i]]--
			X[R[i]]--
		}
	}
}
