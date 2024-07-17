package main

import "fmt"

func main() {
	d, n := 0, 0
	fmt.Scan(&d, &n)
	work := make([]int, d)
	for i := 0; i < d; i++ {
		work[i] = 24
	}

	l, r, h := 0, 0, 0
	for i := 0; i < n; i++ {
		fmt.Scan(&l, &r, &h)
		for j := l - 1; j <= r-1; j++ {
			if work[j] > h {
				work[j] = h
			}
		}
	}

	ans := 0
	for i := 0; i < d; i++ {
		ans += work[i]
	}

	fmt.Println(ans)
}
