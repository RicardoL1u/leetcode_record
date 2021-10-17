package main

import "fmt"

// func getK(int []nums) int {
// 	n := len(nums)
// 	findK(nums, 0, n)

// }

var test = make(map[string]string, 0)

// var test map[string]string

func QuickSort(nums []int, s, e int) {
	if e <= s {
		return
	}
	k := nums[s]
	i, j := s, e
	for i != j {
		for i < j && nums[j] > k {
			j--
		}
		nums[i] = nums[j]
		for i < j && nums[i] < k {
			i++
		}
		nums[j] = nums[i]
	}
	nums[i] = k
	QuickSort(nums, 0, i-1)
	QuickSort(nums, i+1, e)
}

func main() {
	nums := []int{1, 0, 2, 5, 8, 3, 4, 6, 7, 9}
	// test := make(map[string]string, 0)
	names := []byte{'a', 'b', 'c', 'd'}
	fmt.Println(names)
	names1 := string(names)
	fmt.Println(names1)
	test["China"] = "Beijing"
	for k, v := range test {
		fmt.Println(k, v)
	}
	QuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
