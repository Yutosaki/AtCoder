package main

import "fmt"

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}

func main() {
	S := ""
	T := ""
	fmt.Scan(&S, &T)

	longest := make([][]int, len(S)+1)
	for i := 0; i <= len(S); i++ {
		longest[i] = make([]int, len(T)+1)
	}

	for i := 1; i <= len(S); i++ {
		for j := 1; j <= len(T); j++ {
			if S[i-1] == T[j-1] {
				longest[i][j] = longest[i-1][j-1] + 1
			} else {
				longest[i][j] = max(longest[i][j-1], longest[i-1][j])
			}
		}
	}

	fmt.Println(longest[len(S)][len(T)])
}
