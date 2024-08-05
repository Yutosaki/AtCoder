package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	answer := ""
	for i := 1; i <= 10; i++ {
		answer = fmt.Sprintf("%d", n%2) + answer
		n = n / 2
	}
	fmt.Println(answer)
}
