package main

import "fmt"

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	ans := make([]byte, len(s))
	dp := make([][]int, len(s))
	maxLen := 1
	ans = []byte(s[0:1])
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 1
		if i >= 1 && s[i] == s[i-1] {
			dp[i-1][i] = 2
			if dp[i-1][i] > maxLen {
				ans = []byte(s[i-1 : i+1])
				maxLen = dp[i-1][i]
			}
		}
	}
	for i := 1; i < len(s); i++ {
		for offset := 1; i+offset < len(s) && i-offset >= 0; offset++ {
			if s[i+offset] == s[i-offset] {
				dp[i-offset][i+offset] =
					dp[i-offset+1][i+offset-1] + 2
				if dp[i-offset][i+offset] > maxLen {
					ans = []byte(s[i-offset : i+offset+1])
					maxLen = dp[i-offset][i+offset]
				}
			}
		}
		for offset := 1; i+offset < len(s) && i-1-offset >= 0; offset++ {
			r, l := i+offset, i-1-offset
			if s[r] == s[l] && s[l+1] == s[r-1] {
				dp[l][r] =
					dp[l+1][r-1] + 2
				if dp[l][r] > maxLen {
					ans = []byte(s[l : r+1])
					maxLen = dp[l][r]
				}
			}
		}
	}

	return string(ans)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {
	// s := "accbddb"
	// s := "ac"
	s := "aacabdkacaa"
	fmt.Printf("%s\n", longestPalindrome(s))

}
