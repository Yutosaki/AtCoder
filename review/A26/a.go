package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	for i := 0; i < n; i++ {
		if isPrime(A[i]) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
