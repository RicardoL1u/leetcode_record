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
	for j := 1; j < len(w2); j++ {
		isMatched := false
		for i := j; i < len(w1); i++ {
			if w1[i] == w2[j] {
				dp[i][j] = dp[i-1][j-1]
				isMatched = true
				if j == len(w2)-1 {
					printMat(dp, len(w1))
					return dp[i][j]
				}
			} else {
				if isMatched {
					dp[i][j] = dp[i-1][j] + 1 // delete
					continue
				}
				dp[i][j] = dp[i-1][j-1] + 1 // replace
				isMatched = true
				if j == len(w2)-1 {
					printMat(dp, len(w1))
					return dp[i][j]
				}
			}
		}
	}

	return 1
}

func printMat(mat [][]int, row int) {
	for i := 0; i < row; i++ {
		fmt.Println(mat[i])
	}
}

func main() {
	fmt.Println(minDistance("horse", "rose"))
}
