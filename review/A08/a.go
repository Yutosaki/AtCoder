package main

import "fmt"

func main() {
	H, W := 0, 0
	fmt.Scan(&H, &W)

	subset := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		subset[i] = make([]int, W+1)
	}

	tmp := 0
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			fmt.Scan(&tmp)
			subset[i][j] = subset[i][j-1] + tmp
		}
	}

	q := 0
	fmt.Scan(&q)
	a, b, c, d := 0, 0, 0, 0
	ans := make([]int, q)
	for i := 0; i < q; i++ {
		fmt.Scan(&a, &b, &c, &d)
		for j := a; j <= c; j++ {
			ans[i] += subset[j][d] - subset[j][b-1]
		}
	}

	for i := 0; i < q; i++ {
		fmt.Println(ans[i])
	}
}
