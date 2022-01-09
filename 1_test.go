package main

import (
	"log"
	"testing"
)

func twoSum(nums []int, target int) []int {
	rem := make(map[int]int)
	for k, v := range nums {
		rem[target-v] = k
	}
	for k1, v1 := range nums {
		if k2, ok := rem[v1]; ok && k1 != k2 {
			return []int{k1, k2}
		}
	}
	return nil
}

func TestTwoSum(t *testing.T) {
	log.Println(twoSum([]int{3, 2, 4}, 6))
}
