package main

func longestPalindrome(s string) string {
	if len(s) <= 1 {
		return s
	}
	str, end := 0, 1
	maxLen := 1

	for i := 1; i < len(s); i++ {
		for offset := 0; i+offset < len(s) && i-offset >= 0; offset++ {
			l, r := i-offset, i+offset
			if s[r] == s[l] {
				if offset*2+1 > maxLen {
					str, end = l, r+1
					maxLen = offset*2 + 1
				}
			} else {
				break
			}
		}
		for offset := 0; i+offset < len(s) && i-1-offset >= 0; offset++ {
			r, l := i+offset, i-1-offset
			if s[r] == s[l] {
				if offset*2+2 > maxLen {
					str, end = l, r+1
					maxLen = offset*2 + 2
				}
			} else {
				break
			}
		}
	}

	return s[str:end]
}

/*
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

/*
func main() {
	// s := "accbddb"
	// s := "ac"
	s := "aacabdkacaa"
	fmt.Printf("%s\n", longestPalindrome(s))

}

*/
