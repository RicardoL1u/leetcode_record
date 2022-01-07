package main

import "testing"

func findAnagrams(s string, p string) []int {
	pSet := make(map[byte]int)
	for _, v := range p {
		pSet[byte(v)] += 1
	}
	sSet := make(map[byte]int)
	for 
	return nil
}

func compareTwoMap(a, b map[byte]bool) bool {
	for k, _ := range a {
		if !b[k] {
			return false
		}
	}
	for k, _ := range b {
		if !a[k] {
			return false
		}
	}
	return true
}
func TestFindAnagrams(t *testing.T) {
	testMap1 := map[byte]bool{
		'a': true,
	}
	testMap2 := map[byte]bool{
		'a': true,
	}
	// cmpMap := make(map[map[byte]bool]bool)
	if !compareTwoMap(testMap1, testMap2) {
		t.Errorf("why not equal")
	}
}
