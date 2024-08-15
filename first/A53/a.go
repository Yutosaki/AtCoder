package main

import (
	"container/heap"
	"fmt"
)

// IntHeap は整数の最小ヒープです
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	var n int
	fmt.Scan(&n)

	h := &IntHeap{}
	heap.Init(h)

	for i := 0; i < n; i++ {
		var op int
		fmt.Scan(&op)

		switch op {
		case 1:
			var cost int
			fmt.Scan(&cost)
			heap.Push(h, cost)
		case 2:
			if h.Len() > 0 {
				fmt.Println((*h)[0])
			}
		case 3:
			if h.Len() > 0 {
				heap.Pop(h)
			}
		}
	}
}
