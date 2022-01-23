package moni

func firstUniqChar(s string) byte {
	recMap := make(map[rune]int)
	// ans := 0
	for k, v := range s {
		if recMap[v] == 0 {
			recMap[v] = -(k + 1)
		} else {
			recMap[v] = absInt(recMap[v])
		}
	}
	ans := 1 << 30
	for _, v := range recMap {
		if v < 0 && (ans > absInt(v)) {
			ans = absInt(v)
		}
	}
	return s[ans-1]
}

func absInt(a int) int {
	if a > 0 {
		return a
	} else {
		return -a
	}
}
