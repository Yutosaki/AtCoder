package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Scan(&n)

	A := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&A[i])
	}

	// dp配列は最長増加部分列を効率的に計算するためのもの
	dp := []int{}

	for i := 0; i < n; i++ {
		// dpに対してA[i]を挿入する場所を二分探索で見つける
		pos := sort.Search(len(dp), func(j int) bool { return dp[j] >= A[i] })
		
		// もしposがdpの長さより小さければ、そこを更新
		if pos < len(dp) {
			dp[pos] = A[i]//これは実際の部分列ではなく中間的な要素である。dp配列は最長増加列のの末端の最小値を保持している
		} else {
			// それ以外の場合は新しい要素を追加
			dp = append(dp, A[i])
		}
	}

	fmt.Println(len(dp))
}
