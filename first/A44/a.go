package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	readInts := func() int {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		return num
	}

	n, q := readInts(), readInts()

	array := make([]int, n)
	for i := 0; i < n; i++ {
		array[i] = i + 1
	}

	state := 1

	for i := 0; i < q; i++ {
		first := readInts()
		if first == 1 { //change
			before, after := readInts(), readInts()
			//xを探してyにする
			if state == 1 {
				array[before-1] = after
			} else {
				array[n-1-(before-1)] = after
			}
		} else if first == 2 { //reverse
			if state == 1 {
				state = 2
			} else {
				state = 1
			}
		} else {
			which := readInts()
			if state == 1 {
				fmt.Println(array[which-1])
			} else {
				fmt.Println(array[n-1-(which-1)])
			}
		}
	}
}
