package main

import "fmt"

func longestValidParentheses1(s string) int {
	dp := make([][]int, len(s)+2)
	for k := range dp {
		dp[k] = make([]int, len(s)+2)
	}
	ans := 0
	for i := 1; i <= len(s); i++ {
		for j := i; j <= len(s); j++ {
			if s[j-1] == '(' {
				if dp[i][j-1] >= 0 {
					dp[i][j] = dp[i][j-1] + 1
				} else {
					break
				}
			} else {
				dp[i][j] = dp[i][j-1] - 1
			}
			if dp[i][j] == 0 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

func longestValidParentheses(s string) int {
	ans := 0
	temp := 0
	for i := 1; i <= len(s); i++ {
		temp = 0
		for j := i; j <= len(s); j++ {
			if s[j-1] == '(' {
				if temp >= 0 {
					temp++
				} else {
					break
				}
			} else {
				temp--
			}
			if temp == 0 {
				ans = max(ans, j-i+1)
			}
		}
		// fmt.Println(dp)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}
func main() {
	// s := ")()())"
	fmt.Println(longestValidParentheses(")()())") == 4)
	fmt.Println(longestValidParentheses(")()(()))") == 6)
	fmt.Println(longestValidParentheses(")") == 0)
	fmt.Println(longestValidParentheses("") == 0)
	fmt.Println(longestValidParentheses(")))))(((((((") == 0)
	fmt.Println(longestValidParentheses(")))))((((((())))") == 8)
	fmt.Println(longestValidParentheses("()") == 2)
}
