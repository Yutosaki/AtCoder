package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	X := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&X[i])
	}

	answer := make([]bool, n)
	for i := 0; i < n; i++ {
		answer[i] = isPrime(X[i])
	}

	for i := 0; i < n; i++ {
		if answer[i] == true {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(num int) bool {
	if num == 2 || num == 3 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	for i := 3; i <= num/i; i += 2 {
		if num%i == 0 {
			return false
		}
	}
	return true
}
