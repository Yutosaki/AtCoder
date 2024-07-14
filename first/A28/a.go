package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	answer := make([]int, n+1)
	a := 0

	for i := 1; i <= n; i++ {
		operator := ""
		fmt.Scan(&operator, &a)
		if operator == "+" {
			answer[i] = answer[i-1] + a
		} else if operator == "-" {
			answer[i] = answer[i-1] - a
		} else {
			answer[i] = answer[i-1] * a
		}
		answer[i] %= 10000
		if answer[i] < 0 {
			answer[i] += 10000
		}
	}

	for i := 1; i <= n; i++ {
		fmt.Println(answer[i])
	}
}
