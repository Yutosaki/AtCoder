package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	arrayA := make([]int, n+1)
	arrayB := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Scan(&arrayA[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Scan(&arrayB[i])
	}

	timer := make([]int, n+1)
	timer[2] = arrayA[2]
	for i := 3; i <= n; i++ {
		if timer[i-2]+arrayB[i] < timer[i-1]+arrayA[i] {
			timer[i] = timer[i-2] + arrayB[i]
		} else {
			timer[i] = timer[i-1] + arrayA[i]
		}
	}

	fmt.Println(timer[n])
}
