package main

import "fmt"

func main() {
	n, r := 0, 0
	fmt.Scan(&n, &r)
	mod := 1000000007
	up := 1
	for i := 1; i <= n; i++ {
		up = (up * i) % mod
	}
	down := 1
	for i := 1; i <= r; i++ {
		down = (down * i) % mod
	}
	for i := 1; i <= n-r; i++ {
		down = (down * i) % mod
	}

	fmt.Println(up * power(down, mod-2, mod) % mod)
}

func power(base, exponent, mod int) int {
	result := 1
	p := base
	for i := 0; i < 30; i++ {
		bitCheck := 1 << i
		if (exponent/bitCheck)%2 == 1 {
			result = (result * p) % mod
		}
		p = (p * p) % mod
	}
	return result
}
