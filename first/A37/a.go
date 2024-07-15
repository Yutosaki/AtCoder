package main

import "fmt"

func main() {
	n, m, b := 0, 0, 0
	fmt.Scan(&n, &m, &b)
	sumA := 0
	for i := 0; i < n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		sumA += tmp
	}
	sumC := 0
	for i := 0; i < m; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		sumC += tmp
	}

	fmt.Println(n*m*b + m*sumA + n*sumC)
}
