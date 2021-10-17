package main

import (
	"fmt"
)

func findRepeatedDnaSequences(s string) []string {
	ans := make([]string, 0)
	ansSet := make(map[string]bool, 0)
	if len(s) <= 10 {
		return ans
	} else {
		// ans = append(ans, s)
	}
	temp := s[0:10]
	for i := 0; i < len(s)-10; i++ {
		temp = s[i : i+10]
		if ansSet[temp] == true {
			continue
		}
		// fmt.Print("hi")
		for j := i + 1; j <= len(s)-10; j++ {
			rep := s[j : j+10]
			fmt.Print(rep)
			if rep == temp {
				// ans = append(ans, temp)
				ansSet[temp] = true
				break
			}
		}
	}
	for k, _ := range ansSet {
		ans = append(ans, k)
	}
	return ans
}

func main() {
	// test := "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
	test := "AAAAAAAAAAA"
	fmt.Print(findRepeatedDnaSequences(test))

}
