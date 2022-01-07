package main

import (
	"fmt"
	"strings"
)

func arraySatisfy(cnt, req []int) bool {
	mark := true
	for k, v := range cnt {
		if req[k] > v {
			return !mark
		}
	}
	return mark
}

func maxInt(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// var stop int = 0

func minWindow(s, t string) string {
	str, end := getAns(s, t)
	resPoi := 0
	minn := 1 << 30
	fmt.Println(str, end)
	if len(end) == 0 {
		return ""
	} else {
		for k := range end {
			if end[k]-str[k] < minn {
				minn = end[k] - str[k]
				resPoi = k
			}
		}
		ansLeft := str[resPoi]
		ansRight := end[resPoi]
		res := s[ansLeft : ansRight+1]
		return res
	}
}

func getAns(s string, t string) ([]int, []int) {
	str := []int{}
	end := []int{}
	l := 0
	r := 0
	n := len(s)
	req := make([]int, 255)
	cnt := make([]int, 255)
	for _, v := range t {
		req[v]++
	}
	for l < n {
		if strings.ContainsRune(t, rune(s[l])) {
			break
		}
		l++
	}
	r = l
	for l < n && r < n {
		// fmt.Println(l, r, cnt[65:70], str, end)
		str = append(str, l)
		cnt[s[l]]++

		if arraySatisfy(cnt, req) {
			end = append(end, r)
			cnt[s[l]]--
			l++
			// fmt.Println(l)
			for l < n {
				cnt[s[l]]--
				if strings.ContainsRune(t, rune(s[l])) {
					break
				}
				l++
			}
			continue
		}

		r = maxInt(r+1, l+1)
		for r < n {
			cnt[s[r]]++
			if arraySatisfy(cnt, req) {
				end = append(end, r)
				cnt[s[l]]--
				l++
				for l < n {
					cnt[s[l]]--
					if strings.ContainsRune(t, rune(s[l])) {
						break
					}
					l++
				}
				break
			}
			r++
		}

	}
	return str, end
}

/*
func main() {
	s := "ADOBECODEBABC"
	t := "ABC"
	fmt.Println(minWindow(s, t))
	fmt.Println(minWindow("zzz", "zzz"))
	fmt.Println(minWindow("ab", "b"))
	fmt.Println(minWindow("cabwefgewcwaefgcf", "cae"))

}

*/
