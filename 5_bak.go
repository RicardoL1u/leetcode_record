package main

import "fmt"

func longestPalindromebak(s string) string {
	if len(s) <= 1 {
		return s
	}
	rev := make([]byte, 0)
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, s[i])
	}
	max_len := 0
	now_len := 0
	ans := make([]byte, len(s))
	for offset := -(len(s) - 1); offset < len(s)-1; offset++ {
		p1, p2 := calptr(offset)
		now_len = 0
		fmt.Println()
		fmt.Println(offset)
		for p1 < len(s) && p2 < len(s) {
			fmt.Println(s[p1], rev[p2])
			if s[p1] == rev[p2] {
				now_len++
				max_len = max(now_len, max_len)
				if (p1 == len(s)-1 || p2 == len(s)-1) && now_len == max_len {
					if p2-now_len < 0 {
						ans = []byte(s[p1-now_len+1 : p1+1])
					} else {
						ans = []byte(rev[p2-now_len+1 : p2+1])
					}
				}
			} else {
				if now_len == max_len {
					ans = []byte(rev[p2-now_len : p2])
					now_len = 0
				}
			}
			fmt.Print(string(ans))
			p1++
			p2++
		}
	}
	return string(ans)
}

func calptr(offset int) (p1, p2 int) {
	if offset >= 0 {
		p1 = offset
		p2 = 0
	} else {
		p1 = 0
		p2 = -offset
	}
	return
}
func longestPalindrome1(s string) string {
	if len(s) <= 1 {
		return s
	}
	ans := make([]byte, 1000)
	max_len := 0
	p1, p2 := 0, len(s)-1
	for p1 < len(s) {
		p2 = len(s) - 1
		for p2 >= p1 {
			if p2-p1+1 > max_len && isRep(p1, p2, s) {
				max_len = p2 - p1 + 1
				ans = []byte(s[p1 : p2+1])
				break
			}
			p2--
		}
		p1++
	}
	return string(ans)
}

// /*
// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }

func isRep(p1, p2 int, s string) bool {
	for p1 <= p2 {
		if s[p1] != s[p2] {
			return false
		}
		p1++
		p2--
	}
	return true
}

/*
func main() {
	// s := "ba"
	// s := "accccba"
	s := "aacabdkacaa"
	fmt.Printf("%s\n", longestPalindrome(s))
	// fmt.Print(isRep(0, len(s)-1, s))
}

*/
