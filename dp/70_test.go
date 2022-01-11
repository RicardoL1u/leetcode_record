package dp

import (
	"log"
	"testing"
)

// dp
// dp[i] = dp[i-1] + dp[i-2]
func climbStairs(n int) int {
	if n <= 1 {
		return n
	}
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1

	for i := 2; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func TestClimb(t *testing.T) {
	log.Println(climbStairs(3))
}
