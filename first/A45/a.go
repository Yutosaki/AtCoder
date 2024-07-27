package main

import (
	"fmt"
)

func main() {
	var N int
	var C string
	fmt.Scan(&N, &C)

	fmt.Println(N, C)
	A := make([]rune, N+1)
	for i := 1; i <= N; i++ {
		fmt.Scanf("%c", &A[i])
	}

	score := 0
	for i := 1; i <= N; i++ {
		if A[i] == 'W' {
			score += 0
		}
		if A[i] == 'B' {
			score += 1
		}
		if A[i] == 'R' {
			score += 2
		}
	}

	if score%3 == 0 && C == "W" {
		fmt.Println("Yes")
	} else if (score%3 == 1 || score%3 == -2) && C == "B" {
		fmt.Println("Yes")
	} else if (score%3 == 2 || score%3 == -1) && C == "R" {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
