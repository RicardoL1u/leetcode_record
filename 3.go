package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	if len(s) <= 1 {
		return len(s)
	}
	p1, p2 := 0, 1
	recMap := make(map[byte]int)
	for ch := range s {
		recMap[s[ch]] = -1
	}
	recMap[s[p1]] = 0
	ans := 1
	for p1 < len(s) && p2 < len(s) {
		if recMap[s[p2]] < p1 {
			recMap[s[p2]] = p2
		} else {
			p1 = recMap[s[p2]] + 1
			recMap[s[p2]] = p2
		}
		p2++
		ans = max_int(p2-p1, ans)
	}
	return ans
}
func max_int(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func main() {
	// s := "aaaabcd"
	// s := "pwwkew"
	// s := "pae cd c"
	s := "aab"
	fmt.Printf("%d\n", lengthOfLongestSubstring(s))
}
