package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)

	stack := make([]string, 0)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		if num == 1 {
			title := ""
			fmt.Scan(&title)
			stack = append(stack, title)
		} else if num == 2 {
			fmt.Printf("%s\n", stack[len(stack)-1])
		} else if num == 3 {
			stack = stack[:len(stack)-1]
		}
	}
}
