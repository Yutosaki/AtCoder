package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	n := 0
	fmt.Scan(&n)

	A := make([]int, n+1)
	B := make([]int, n+1)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInt := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}

	for i := 2; i <= n; i++ {
		A[i] = readInt()
	}
	for i := 3; i <= n; i++ {
		B[i] = readInt()
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

	fmt.Println(root[n])
}
