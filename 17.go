package main

import "fmt"

var numMap = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	getUnitCmb(0, digits, "", &ans)
	return ans
}

func getUnitCmb(now int, digits, temp string, ans *[]string) {
	if now == len(digits) {
		// fmt.Println(temp)
		if temp != "" {
			*ans = append(*ans, temp)
		}
		return
	}
	for _, ch := range numMap[byte(digits[now])] {
		getUnitCmb(now+1, digits, temp+string(ch), ans)
	}
}

func testIntSlice(test []int) {
	addr0 := &test
	test[0] = 2
	test = append(test, 123)
	addr := &test
	fmt.Printf("%x\n", addr)
	fmt.Printf("%x\n", addr0)

}

/*
func main() {
	test := []int{
		1, 2, 3, 4,
	}
	testIntSlice(test)
	addr := &test
	fmt.Printf("%x\n", addr)
	fmt.Println(test)
	// fmt.Println(letterCombinations(""))
}

*/
