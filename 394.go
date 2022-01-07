package main

func decodeString(s string) string {
	p := 0
	return realDecodeString(s, &p)
}

func realDecodeString(s string, p *int) string {
	ans := ""
	for *p < len(s) {
		if s[*p] <= '9' && s[*p] >= '0' {
			ans += formatRepeatStr(s, p)
		} else if s[*p] != ']' {
			ans += string(s[*p])
		} else {
			return ans
		}
		*p++
	}
	return ans
}

func formatRepeatStr(s string, p *int) string {
	temp := 0
	for s[*p] <= '9' && s[*p] >= '0' {
		temp = temp*10 + int(s[*p]-'0')
		*p++
	}
	*p++ // omit '['
	tempStr := realDecodeString(s, p)
	ret := ""
	for i := 0; i < temp; i++ {
		ret += tempStr
	}
	return ret
}

/*
func main() {
	log.Println(decodeString("abc3[cd]xyz") == "abccdcdcdxyz")
	log.Println(decodeString("3[a]2[bc]") == "aaabcbc")
	log.Println(decodeString("3[a2[c]]") == "accaccacc")
	// log.
}

*/
