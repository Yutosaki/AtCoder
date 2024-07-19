package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	ans := ""
	for i := 0; i < 10; i++ {
		ans = fmt.Sprintf("%d", n%2) + ans
		n /= 2
	}

	fmt.Println(ans)
}
