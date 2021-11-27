package main

import "fmt"

func isMatch(s string, p string) bool {
	a, b := 0, 0
	p1, p2 := &a, &b
	return tryMatch(p1, p2, s, p)
}

func tryMatch(str1, str2 *int, s, p string) bool {
	fmt.Println(*str1, *str2)
	if *str1 == len(s) || *str2 == len(p) {
		if *str1 == len(s) && *str2 == len(p) {
			fmt.Println("hi")
			return true
		} else {
			return false
		}
	}
	switch p[*str2] {
	case '*':
		return matchStar(str1, str2, s, p)
	case '.':
		*str1++
		*str2++
		return tryMatch(str1, str2, s, p)
	default:
		if s[*str1] == p[*str2] {
			*str1++
			*str2++
			return tryMatch(str1, str2, s, p)
		}
	}
	return true
}

func matchStar(str1, str2 *int, s, p string) bool {
	for s[*str1] == p[*str2-1] || p[*str2-1] == '.' {
		*str1++
	}
	*str2++
	return tryMatch(str1, str2, s, p)
}

func main() {
	// s := "aab"
	// p := "c*a*b"
	s := "mississippi"
	p := "mis*is*p*."
	fmt.Println(isMatch(s, p))
}
