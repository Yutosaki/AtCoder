package main

import "fmt"

func main() {
	n, a, b := 0, 0, 0
	fmt.Scan(&n, &a, &b)

	//firstがその数に至るのは可能かどうかで判別
	dp := make([]bool, n+1)
	for i := 1; i <= n; i++ {
		if i-a >= 0 && dp[i-a] == false {
			dp[i] = true
		} else if i-b >= 0 && dp[i-b] == false {
			dp[i] = true
		} else {
			dp[i] = false
		}
	}
	if dp[n] == true {
		fmt.Println("First")
	} else {
		fmt.Println("Second")
	}
}
