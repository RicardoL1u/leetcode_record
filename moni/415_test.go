package moni

import (
	"log"
	"strconv"
	"testing"
)

func addStrings(num1 string, num2 string) string {
	num1, num2 = paddingStr(num1, num2)
	// log.Println(num1, num2)
	ans := ""
	note, now := 0, 0
	for i := len(num1) - 1; i >= 0; i-- {
		now = int(num2[i]+num1[i]-'0'-'0') + note
		note = now / 10
		ans = strconv.Itoa(now%10) + ans
	}
	if note > 0 {
		ans = strconv.Itoa(note%10) + ans
	}
	return ans
}

func paddingStr(num1, num2 string) (string, string) {
	diff := len(num1) - len(num2)
	if diff > 0 {
		for i := 0; i < diff; i++ {
			num2 = "0" + num2
		}
	} else if diff < 0 {
		diff = -diff
		for i := 0; i < diff; i++ {
			num1 = "0" + num1
		}
	} else {
		return num1, num2
	}
	return num1, num2
}

func Test415(t *testing.T) {
	log.Println(addStrings("111", "34") == "145")
	log.Println(addStrings("111", "999") == "1110")
}
