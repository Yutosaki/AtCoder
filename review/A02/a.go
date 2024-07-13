package main

import "fmt"

func main() {
	n, x := 0, 0
	fmt.Scan(&n, &x)
	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}
	for i := 0; i < n; i++ {
		if A[i] == x {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
