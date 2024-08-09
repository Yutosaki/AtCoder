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

	revRoot := make([]int, 1)
	revRoot[0] = n
	for i := n; i >= 1; i-- {
		if root[i] == root[i-1]+A[i] {
			if i-1 != revRoot[len(revRoot)-1] {
				revRoot = append(revRoot, i-1)
			}
		} else {
			if i-2 != revRoot[len(revRoot)-1] {
				revRoot = append(revRoot, i-2)
			}
			i--
		}
	}

	fmt.Println(len(revRoot) - 1)
	for i := len(revRoot) - 2; i >= 1; i-- {
		fmt.Printf("%d ", revRoot[i])
	}
	fmt.Println(n)
}
