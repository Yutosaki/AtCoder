package main

import "fmt"

func main() {
	n, k := 0, 0
	fmt.Scan(&n, &k)

	if k < n*2-2 {
		fmt.Println("No")
		return
	}
	if (k-(n*2-2))%2 == 1 {
		fmt.Println("No")
		return
	}
	fmt.Println("Yes")
}
