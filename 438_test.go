package main

import (
	"log"
	"testing"
)

func findAnagrams(s string, p string) []int {
	ans := make([]int, 0)
	diff := make(map[byte]int)
	diffTot := 0
	for _, v := range p {
		diff[byte(v)] += 1
		diffTot++
	}
	for k, v := range s {
		diff[byte(v)] -= 1
		if diff[byte(v)] >= 0 {
			diffTot--
		} else {
			diffTot++
		}
		if k >= len(p) {
			toRMByte := byte(s[k-len(p)])
			diff[toRMByte] += 1
			if diff[toRMByte] <= 0 {
				diffTot--
			} else {
				diffTot++
			}
		}
		if k >= len(p)-1 && diffTot == 0 {
			ans = append(ans, k-len(p)+1)
		}
	}
	return ans
}

func TestFindAnagrams(t *testing.T) {
	log.Println(findAnagrams("cbaebabacd", "abc"))
}

func TestMap(t *testing.T) {
	testMap1 := map[byte]int{
		'a': 1,
	}
	// testMap2 := map[byte]int{
	// 	'a': 1,
	// 	'b': 2,
	// }
	delete(testMap1, 'a')
	log.Println(len(testMap1))
	testMap1['a'] = 1
	log.Println(len(testMap1))
}

// func compareTwoMap(a, b map[byte]int) map[byte]int {
// 	diffMap := make(map[byte]int)
// 	mark := true
// 	for k, _ := range a {
// 		if b[k] != a[k] {
// 			mark = false
// 			diffMap[k] += 1
// 		}
// 	}
// 	for k, _ := range b {
// 		if a[k] != b[k] {
// 			mark = false
// 			diffMap[k] -= 1
// 		}
// 	}
// 	return diffMap
// }

// func TestCompareMap(t *testing.T) {
// 	testMap1 := map[byte]int{
// 		'a': 1,
// 	}
// 	testMap2 := map[byte]int{
// 		'a': 1,
// 		'b': 2,
// 	}
// 	// cmpMap := make(map[map[byte]bool]bool)
// 	if len(compareTwoMap(testMap1, testMap2)) == 0 {
// 		t.Errorf("why not equal")
// 	}
// }
