package main

import "fmt"

func main() {
	h, w, n := 0, 0, 0
	fmt.Scan(&h, &w, &n)

	matrix := make([][]int, h+1)
	for i := 0; i <= h; i++ {
		matrix[i] = make([]int, w+1)
	}

	for i := 0; i < n; i++ {
		A, B, C, D := 0, 0, 0, 0
		fmt.Scan(&A, &B, &C, &D)
		matrix[A-1][B-1]++
		matrix[A-1][D]--
		matrix[C][B-1]--
		matrix[C][D]++
	}

	for i := 0; i <= h; i++ {
		for j := 1; j <= w; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	for j := 0; j <= w; j++ {
		for i := 1; i <= h; i++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			fmt.Print(matrix[i][j])
			if j != w {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
