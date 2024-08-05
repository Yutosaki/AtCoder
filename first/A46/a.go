package main

import "fmt"

func main() {
	n := 0
	fmt.Scan(&n)
	X := make([]int, n+1)
	Y := make([]int, n+1)
	for i := 1; i <= n; i++ {
		fmt.Scan(&X[i], &Y[i])
	}

	visited := make([]bool, n+1)
	currentPlace := 1
	visited[1] = true
	root := make([]int, n+1)
	root[1] = 1
	for i := 2; i <= n; i++ {
		minDistance := 1000000000
		for j := 1; j <= n; j++ {
			if visited[j] {
				continue
			}
			newDistance := calculateDistance(X[currentPlace]-X[j], Y[currentPlace]-Y[j])
			if newDistance < minDistance {
				minDistance = newDistance
				currentPlace = j
			}
		}
		visited[currentPlace] = true
		root[i] = currentPlace
	}

	for i := 1; i <= n; i++ {
		fmt.Println(root[i])
	}
	fmt.Println("1")
}

func calculateDistance(x, y int) int {
	return x*x + y*y
}
