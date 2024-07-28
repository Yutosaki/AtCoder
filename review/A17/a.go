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
			root[i] = A[i] + root[i-1]
		} else {
			root[i] = B[i] + root[i-2]
		}
	}

	answer := make([]int, 0)
	place := n
	answer = append(answer, place)
	for place != 1 {
		if root[place] == root[place-1]+A[place] {
			answer = append(answer, place-1)
			place--
		} else {
			answer = append(answer, place-2)
			place -= 2
		}
	}
	fmt.Println(len(answer))

	for i := len(answer) - 1; i >= 0; i-- {
		fmt.Printf("%d ", answer[i])
	}
	fmt.Println()
}
