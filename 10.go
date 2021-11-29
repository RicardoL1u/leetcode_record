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
	pt := getOpArr("a" + p + "a")
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
					dp[p1][p2] = true
					p1++
					for p1 < len(st) && (st[p1] == pt[p2].opr || pt[p2].opr == '.') {
						dp[p1][p2] = true
						p1++
					}
				}
			} else {
				if st[p1] == pt[p2].opr || pt[p2].opr == '.' {
					dp[p1][p2] = dp[p1][p2] || dp[p1-1][p2-1]
				}
			}
		}
	}
	for j := 0; j < len(pt); j++ {
		for i := 0; i < len(st); i++ {
			fmt.Printf("%v ", dp[i][j])
		}
		fmt.Println()
	}
	return dp[len(st)-1][len(pt)-1]
}

func isMatchSpc(s string, p string) bool {
	st := "a" + s + "a"
	pt := getOpArr("a" + p + "a")
	dp := make([]bool, len(st))
	// for i := 0; i < len(st); i++ {
	// 	dp[i] = make([]bool, len(pt))
	// }
	// dp[0] = true
	for p2 := 1; p2 < len(pt); p2++ {
		now, pre := false, false
		for p1 := 1; p1 < len(st); p1++ {
			if p1 == 1 && p2 == 1 {
				pre = true
			}
			now = dp[p1]
			if pt[p2].isStar && pre {
				// dp[p1-1][p2] = true
				if st[p1] == pt[p2].opr || pt[p2].opr == '.' {
					dp[p1] = true
					p1++
					for p1 < len(st) && (st[p1] == pt[p2].opr || pt[p2].opr == '.') {
						dp[p1] = true
						p1++
					}
				} else {
					dp[p1] = false
				}
			} else {
				if st[p1] == pt[p2].opr || pt[p2].opr == '.' {
					dp[p1] = pre
				}
			}
			pre = now
		}
		fmt.Println(dp)
		dp[0] = false
	}

	return dp[len(st)-1]
}

func getOpArr(p string) (opArr []op) {
	opArr = make([]op, 0)
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
	return
}

func main() {

	// strings.SplitAfter()
	// fmt.Println(getOpArr("m.*ii"))
	fmt.Printf("Test 2 %v\n", false == isMatch("mississippi", "m.*ii"))
	// fmt.Printf("Test 1 %v\n", true == isMatchSpc("mississippi", "m.*i"))
	fmt.Printf("Test 2 %v\n", false == isMatchSpc("mississippi", "m.*ii"))
	// fmt.Printf("Test 3 %v\n", true == isMatchSpc("mississippi", "m.*"))
	// fmt.Printf("Test 4 %v\n", true == isMatchSpc("aaa", "a.*a"))
	// fmt.Printf("Test 5 %v\n", true == isMatchSpc("a", "a*"))
	// fmt.Printf("Test 6 %v\n", true == isMatchSpc("aab", "c*a*b"))
	// fmt.Printf("Test 7 %v\n", true == isMatchSpc("a", "ab*"))
}
