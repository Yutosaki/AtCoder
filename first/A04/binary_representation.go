package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	array := [10]int{}
	for i := 9; i >= 0; i-- {
		array[i] = n % 2
		n /= 2
	}
	for i := 0; i < 10; i++ {
		fmt.Print(array[i])
	}
	fmt.Println()
}
