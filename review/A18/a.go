package main

import "fmt"

func main() {
	n, s := 0, 0
	fmt.Scan(&n, &s)

	A := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	matrix := make([][]bool, n+1)
	for i := 0; i <= n; i++ {
		matrix[i] = make([]bool, s+1)
	}

	matrix[0][0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j <= s; j++ {
			if matrix[i-1][j] || (j-A[i] >= 0 && matrix[i-1][j-A[i]]) {
				matrix[i][j] = true
			}
		}
	}

	if matrix[n][s] {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
