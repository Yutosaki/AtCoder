package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n+1)
	B := make([]int, n+1)

	for i := 2; i <= n; i++ {
		fmt.Scan(&A[i])
	}

	for i := 3; i <= n; i++ {
		fmt.Scan(&B[i])
	}

	root := make([]int, n+1)
	root[2] = A[2]
	for i := 3; i <= n; i++ {
		if root[i-1]+A[i] < root[i-2]+B[i] {
			root[i] = root[i-1] + A[i]
		} else {
			root[i] = root[i-2] + B[i]
		}
	}

	revroot := make([]int, 0)
	for i := n; i >= 1; i-- {
		if root[i] == root[i-1]+A[i] {
			revroot = append(revroot, i-1)
		} else if root[i] == root[i-2]+B[i] {
			revroot = append(revroot, i-2)
			i--
		}
	}

	fmt.Println(len(revroot))

	for i := len(revroot) - 2; i >= 0; i-- {
		fmt.Printf("%d ", revroot[i])
	}
	fmt.Println(n)
}
