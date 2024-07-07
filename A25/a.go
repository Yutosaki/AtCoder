package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	low, col := 0, 0
	fmt.Scan(&low, &col)

	matrix := make([][]string, low)
	dp := make([][]int, low)

	for i := 0; i < low; i++ {
		matrix[i] = make([]string, col)
		dp[i] = make([]int, col)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < low; i++ {
		scanner.Scan()
		line := scanner.Text()
		for j := 0; j < col; j++ {
			matrix[i][j] = string(line[j])
		}
	}

	if matrix[0][0] == "." {
		dp[0][0] = 1
	}

	for i := 1; i < low; i++ {
		if matrix[i][0] == "#" {
			dp[i][0] = 0
		} else {
			dp[i][0] = dp[i-1][0]
		}
	}
	for i := 1; i < col; i++ {
		if matrix[0][i] == "#" {
			dp[0][i] = 0
		} else {
			dp[0][i] = dp[0][i-1]
		}
	}

	for i := 1; i < low; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == "#" {
				dp[i][j] = 0
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	fmt.Println(dp[low-1][col-1])
}
