package main

import (
	"fmt"
)

func findDuplicate(nums []int) int {
	tmp := make([]bool, 1e5)
	for _, v := range nums {
		if tmp[v] {
			return v
		}
		tmp[v] = true
	}
	return 10
}

func main() {
	nums := []int{3, 1, 3, 4, 2}
	fmt.Println(nums)
	fmt.Println(findDuplicate(nums))
}
