package main

import "fmt"

func maxCoins(nums []int) [][]int {
	n := len(nums)
	dp := make([][]int, n+2)
	v := make([]int, n+2)
	v[0], v[n+1] = 1, 1
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := 1; i <= n; i++ {
		v[i] = nums[i-1]
	}

	for i := n - 1; i >= 0; i-- {
		for j := i + 2; j <= n+1; j++ {
			for k := i + 1; k < j; k++ {
				sum := v[i] * v[k] * v[j]  //乘上两个端点
				sum += dp[i][k] + dp[k][j] //这两个dp表示后放
				dp[i][j] = max(sum, dp[i][j])
			}
		}
	}
	return dp
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	nums := []int{3, 1, 5, 8}
	dp := maxCoins(nums)
	for k := 0; k < len(dp); k++ {
		for _, e := range dp[k] {
			fmt.Print(e)
			fmt.Print(" ")
		}
		fmt.Println()
	}
}
