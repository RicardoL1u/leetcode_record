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

func longestValidParentheses2(s string) int {
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

func longestValidParenthesesDPN(s string) int {
	ans := 0
	dp := make([]int, len(s)+1)
	st := "a" + s
	for i := 2; i <= len(s); i++ {
		if st[i] == ')' && st[i-1] == '(' {
			dp[i] = dp[i-2] + 2
		} else if st[i] == ')' && st[i-1] == ')' {
			if st[i-dp[i-1]-1] == '(' {
				dp[i] = dp[i-1] + dp[i-dp[i-1]-2] + 2
			}
		}
		ans = max(dp[i], ans)
	}
	return ans
}
func longestValidParentheses(s string) int {
	ans1 := getAns(s)
	st := make([]byte, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		st[len(s)-1-i] = byte(uint8(s[i]) ^ 0x1)
	}
	ans2 := getAns(string(st))
	// fmt.Println(string(st), ans1, ans2)
	return min(ans1, ans2)
}

func getAns(s string) int {
	ans, cnt := 0, 0
	l, r := 0, 0
	for l < len(s) && r < len(s) {
		if s[r] == '(' {
			cnt++
			r++
		} else {
			if cnt == 0 {
				r++
				l = r
				continue
			} else {
				cnt--
			}
		}
		if cnt == 0 {
			ans = max(ans, r-l+1)
		}
	}
	for l < len(s) {
		if s[l] == '(' {
			cnt--
			l++
		}
		if cnt == 0 {
			ans = max(ans, r-l)
			break
		}
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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}

func main() {
	// s := ")()())"
	fmt.Println(longestValidParentheses("(()(((()"))
	fmt.Println(longestValidParentheses(")()(()") == 2)
	fmt.Println(longestValidParentheses(")()())") == 4)
	fmt.Println(longestValidParentheses(")()(()"))
	fmt.Println(longestValidParentheses(")()(()))") == 6)
	fmt.Println(longestValidParentheses(")") == 0)
	fmt.Println(longestValidParentheses(")(") == 0)
	fmt.Println(longestValidParentheses("") == 0)
	fmt.Println(longestValidParentheses(")))))(((((((") == 0)
	fmt.Println(longestValidParentheses(")))))((((((())))") == 8)
	fmt.Println(longestValidParentheses(")))))((((((())))"))
	fmt.Println(longestValidParentheses("()") == 2)
}
