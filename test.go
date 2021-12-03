package main

import "fmt"

func main() {
	// 判断map中元素是否存在
	// 兼容0
	test := map[int]int{1: 1, 2: 2, 3: 0}
	for i := 0; i < 4; i++ {
		if k, ok := test[i]; ok {
			fmt.Println(k, ok)
		} else {
			fmt.Println(k, ok)
		}
	}
}
