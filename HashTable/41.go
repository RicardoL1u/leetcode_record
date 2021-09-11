package main

import "fmt"

func firstMissingPositive(nums []int) int {
	n := len(nums)
	for k := range nums {
		if nums[k] <= 0 {
			nums[k] = 6e6
		}
	}

	for _, v := range nums {
		if v == 0 {
			continue
		} else {

			k := absInt(v) - 1
			if k >= n {
				continue
			} else {
				if nums[k] < 0 {
					continue
				} else {
					nums[k] = -nums[k]
				}
			}
		}
	}

	for k, v := range nums {
		// fmt.Println(v)
		if v > 0 {
			return k + 1
		}
	}
	return len(nums) + 1
}

func absInt(v int) int {
	if v < 0 {
		return -v
	} else {
		return v
	}
}
func main() {
	// nums := []int{1, 2, 3, 5}
	nums := []int{1, 2, 0}
	nums1 := []int{0, 8, 9, 10, 11}
	nums2 := []int{0, 8, 9, 10, 1}
	fmt.Println(firstMissingPositive(nums))
	fmt.Println(firstMissingPositive(nums1))
	fmt.Println(firstMissingPositive(nums2))
}
