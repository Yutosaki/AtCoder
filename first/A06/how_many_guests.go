package main

import "fmt"

func main() {
	n, q := 0, 0
	fmt.Scan(&n, &q)
	subset := make([]int, n+1)

	for i := 1; i <= n; i++ {
		tmp := 0
		fmt.Scan(&tmp)
		subset[i] += subset[i-1] + tmp
	}

	res := make([]int, q)
	for i := 0; i < q; i++ {
		l, r := 0, 0
		fmt.Scan(&l, &r)
		res[i] = subset[r] - subset[l-1]
	}
	for i := 0; i < q; i++ {
		fmt.Println(res[i])
	}
}
