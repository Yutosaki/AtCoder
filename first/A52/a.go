package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	queue := make([]string, 0)
	for i := 0; i < n; i++ {
		num := 0
		fmt.Scan(&num)
		if num == 1 {
			name := ""
			fmt.Scan(&name)
			queue = append(queue, name)
		} else if num == 2 {
			fmt.Println(queue[0])
		} else {
			queue = queue[1:]
		}
	}
}
