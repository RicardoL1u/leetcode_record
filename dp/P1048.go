package main

import (
	"fmt"
)

const MAX = 1010

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	T, M := 0, 0
	fmt.Scanf("%d %d", &T, &M)
	c := make([]int, MAX)
	v := make([]int, MAX)
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d", &c[i], &v[i])
	}
	dp := make([][]int, 110)
	for i := 0; i < 110; i++ {
		dp[i] = make([]int, MAX)
	}
	for i := 1; i <= M; i++ {
		for j := 0; j <= T; j++ {
			if j-c[i-1] >= 0 {
				dp[i][j] = max(dp[i-1][j-c[i-1]]+v[i-1], dp[i-1][j])
				// fmt.Println(i, j, dp[i][j])
			}
		}
	}

	fmt.Println(dp[M][T])

}
