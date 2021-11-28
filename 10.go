package main

import (
	"fmt"
	"strings"
)

type op struct {
	isStar bool
	opr    byte
}

func isMatch(s string, p string) bool {
	st := "a" + s + "a"
	pt := getOpArr(p)
	dp := make([][]bool, len(st))
	for i := 0; i < len(st); i++ {
		dp[i] = make([]bool, len(pt))
	}
	dp[0][0] = true
	for p2 := 1; p2 < len(pt); p2++ {
		for p1 := 1; p1 < len(st); p1++ {
			if pt[p2].isStar && dp[p1-1][p2-1] {
				dp[p1-1][p2] = true
				if st[p1] == pt[p2].opr || pt[p2].opr == '.' {
					dp[p1][p2] = dp[p1][p2] || true && dp[p1-1][p2-1]
					p1++
					for p1 < len(st) && (st[p1] == pt[p2].opr || pt[p2].opr == '.') {
						dp[p1][p2] = dp[p1][p2] || true && dp[p1-1][p2]
						p1++
					}
				}
			} else {
				if st[p1] == pt[p2].opr || pt[p2].opr == '.' {
					dp[p1][p2] = dp[p1][p2] || true && dp[p1-1][p2-1]
				}
			}
		}
	}

	return dp[len(st)-1][len(pt)-1]
}

func getOpArr(p string) (opArr []op) {
	opArr = make([]op, 0)
	opArr = append(opArr, op{isStar: false, opr: 'a'})
	for _, str := range strings.SplitAfter(p, "*") {
		for k, v := range str {
			if k == len(str)-2 && str[len(str)-1] == '*' {
				opArr = append(opArr, op{isStar: true, opr: byte(v)})
				break
			} else {
				opArr = append(opArr, op{isStar: false, opr: byte(v)})
			}
		}
	}
	opArr = append(opArr, op{isStar: false, opr: 'a'})
	return
}

func main() {

	// strings.SplitAfter()
	// fmt.Println(getOpArr("m.*ii"))
	fmt.Printf("Test 1 %v\n", true == isMatch("mississippi", "m.*i"))
	fmt.Printf("Test 2 %v\n", false == isMatch("mississippi", "m.*ii"))
	fmt.Printf("Test 3 %v\n", true == isMatch("mississippi", "m.*"))
	fmt.Printf("Test 4 %v\n", true == isMatch("aaa", "a.*a"))
	fmt.Printf("Test 5 %v\n", true == isMatch("a", "a*"))
	fmt.Printf("Test 6 %v\n", true == isMatch("aab", "c*a*b"))
	fmt.Printf("Test 7 %v\n", true == isMatch("a", "ab*"))
}
