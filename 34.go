package main

import (
	"fmt"
)

func searchRange(nums []int, target int) []int {
	res := []int{0, 0}
	b := upperBound(nums, target)
	if b == -1 {
		res[0] = -1
		res[1] = -1
		return res
	} else {
		a := lowerBound(nums, target)
		res[0] = a
		res[1] = b
		return res
	}
	return res
}

var cnt int = 0

func lowerBound(nums []int, target int) int {
	s := 0
	e := len(nums)

	for s < e {

		mid := (e-s)/2 + s
		// fmt.Println(s, " ", e, " ", mid)
		if nums[mid] < target {
			s = mid + 1
		} else {
			e = mid
		}
	}
	return s
}

func upperBound(nums []int, target int) int {
	s := 0
	e := len(nums)

	for s < e {
		mid := (e-s)/2 + s
		fmt.Println(s, " ", e, " ", mid)
		if nums[mid] <= target {
			s = mid + 1
		} else {
			e = mid
		}
	}
	// fmt.Println(s, " ", e, " ", (e+s)/2)
	if s <= 0 || nums[s-1] != target {
		return -1
	}
	return s - 1
}

func main() {
	nums := []int{0, 1, 1, 2, 2, 3, 3}
	fmt.Println(upperBound(nums, 0))
	fmt.Println(searchRange(nums, 0))
}
