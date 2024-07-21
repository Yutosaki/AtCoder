package main

import "fmt"

func main() {
	h, w, n := 0, 0, 0
	fmt.Scan(&h, &w, &n)

	matrix := make([][]int, h+2)
	for i := 0; i <= h+1; i++ {
		matrix[i] = make([]int, w+2)
	}

	for i := 0; i < n; i++ {
		a, b, c, d := 0, 0, 0, 0
		fmt.Scan(&a, &b, &c, &d)
		matrix[a][b]++
		matrix[a][d+1]--
		matrix[c+1][b]--
		matrix[c+1][d+1]++
	}

	for i := 1; i <= h+1; i++ {
		for j := 1; j <= w+1; j++ {
			matrix[i][j] += matrix[i][j-1]
		}
	}

	for j := 1; j <= w+1; j++ {
		for i := 1; i <= h+1; i++ {
			matrix[i][j] += matrix[i-1][j]
		}
	}

	for i := 1; i <= h; i++ {
		for j := 1; j <= w; j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
}
