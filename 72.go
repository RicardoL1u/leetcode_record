package main

import "fmt"

func minDistance(word1 string, word2 string) int {
	w1, w2 := "a"+word1, "a"+word2
	dp := make([][]int, len(w1))
	for i := 0; i < len(w1); i++ {
		dp[i] = make([]int, len(w2))
	}
	for i := 0; i < len(w1); i++ {
		dp[i][0] = i
	}
	for i := 0; i < len(w2); i++ {
		dp[0][i] = i
	}
	for j := 1; j < len(w2); j++ {
		for i := 1; i < len(w1); i++ {
			if w1[i] == w2[j] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = dp[i-1][j-1] + 1 // replace
			}
			// anyway, we are always able to delete w1[i]
			// to make the dp[i][j] matched
			dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			// same with above, we delete w2[j]
			dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
		}
	}
	// printMat(dp, len(w1))
	return dp[len(w1)-1][len(w2)-1]
}

func min(a, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func printMat(mat [][]int, row int) {
	for i := 0; i < row; i++ {
		fmt.Println(mat[row-i-1])
	}
}

func main() {
	fmt.Println(minDistance("horse", "rose"))
}
