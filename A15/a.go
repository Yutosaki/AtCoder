package main

import (
	"fmt"
	"sort"
)

type numbers struct {
	before  int
	index   int
	pressed int
}

func main() {
	n := 0
	fmt.Scan(&n)
	array := make([]numbers, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&array[i].before)
		array[i].index = i
	}

	sort.Slice(array, func(i, j int) bool { return array[i].before < array[j].before })

	pressed_idx := 0
	for i := 1; i <= n; i++ {
		if array[i].before != array[i-1].before {
			pressed_idx++
		}
		array[i].pressed = pressed_idx
	}

	sort.Slice(array, func(i, j int) bool { return array[i].index < array[j].index })

	for i := 1; i <= n; i++ {
		fmt.Print(array[i].pressed, " ")
	}
	fmt.Println()
}
