package main

import (
	"fmt"
	"sort"
)

func main() {
	n := 0
	fmt.Scan(&n)
	arrayA := make([]int, n+1)
	arrayB := make([]int, n+1)
	for i := 2; i <= n; i++ {
		fmt.Scan(&arrayA[i])
	}
	for i := 3; i <= n; i++ {
		fmt.Scan(&arrayB[i])
	}

	timer := make([]int, n+1)

	timer[2] = arrayA[2]
	for i := 3; i <= n; i++ {
		if timer[i-2]+arrayB[i] < timer[i-1]+arrayA[i] {
			timer[i] = timer[i-2] + arrayB[i]
		} else {
			timer[i] = timer[i-1] + arrayA[i]
		}
	}

	root := make([]int, 1)
	now_place := n
	root[0] = n

	for {
		if now_place == 1 {
			break
		}
		if timer[now_place] == timer[now_place-1]+arrayA[now_place] {
			root = append(root, now_place-1)
			now_place = now_place - 1
		} else {
			root = append(root, now_place-2)
			now_place -= 2
		}
	}

	sort.Slice(root, func(i, j int) bool { return root[i] < root[j] })

	fmt.Println(len(root))
	for i := 0; i < len(root); i++ {
		fmt.Print(root[i], " ")
	}
	fmt.Println()
}
