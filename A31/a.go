package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	three := n / 3
	five := n / 5
	fifeteen := n / 15
	fmt.Println(three + five - fifeteen)
}
