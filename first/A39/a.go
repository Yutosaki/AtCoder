package main

import (
	"fmt"
	"sort"
)

type movie struct {
	start int
	end   int
}

func main() {
	n := 0
	fmt.Scan(&n)
	movies := make([]movie, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&movies[i].start, &movies[i].end)
	}

	sort.Slice(movies, func(i, j int) bool { return movies[i].end < movies[j].end })

	ans := 0
	current := 0
	for i := 0; i < n; i++ {
		if movies[i].start >= current {
			current = movies[i].end
			ans++
		}
	}

	fmt.Println(ans)
}
