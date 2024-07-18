package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	s := make([]rune, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scanf("%c", &s[i])
	}

	flag := false
	for i := 0; i <= n-2; i++ {
		if s[i] == 'R' && s[i+1] == 'R' && s[i+2] == 'R' {
			flag = true
		}
		if s[i] == 'B' && s[i+1] == 'B' && s[i+2] == 'B' {
			flag = true
		}
	}
	if flag == true {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
