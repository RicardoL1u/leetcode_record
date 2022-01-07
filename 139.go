package main

import "fmt"

func wordBreak1(s string, wordDict []string) bool {
	w := s + "#"
	// 由于你定义的是dp[i]表示i个之前的全部匹配
	// 所以只需在后面补 “#”
	dp := make([]bool, len(w))
	dp[0] = true
	i := 0
	for i < len(w) {
		for _, word := range wordDict {
			if i >= len(word) { //同样的由于你的定义与golang的惯用偏移一致，这里只需要对i<len(word)取反就行
				dp[i] = true && dp[i-len(word)] && string(w[i-len(word):i]) == word || dp[i]
			}
		}
		i++
	}
	fmt.Println(dp)
	return dp[len(w)-1]
}
func wordBreak(s string, wordDict []string) bool {
	w := "#" + s
	// 由于你定义的是dp[i]表示i个之前的全部匹配
	// 所以只需在后面补 “#”
	dp := make([]bool, len(w))
	dp[0] = true
	i := 0
	for i <= len(w) {
		for _, word := range wordDict {
			if i > len(word) { //同样的由于你的定义与golang的惯用偏移一致，这里只需要对i<len(word)取反就行
				dp[i-1] = true && dp[i-len(word)-1] && string(w[i-len(word):i]) == word || dp[i-1]
			}
		}
		i++
	}
	fmt.Println(dp)
	return dp[len(w)-1]
}

/*
func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}

*/
