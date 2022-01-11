package luangao

import (
	"log"
	"testing"
)

func moveZeroes(nums []int) {
	// p1, p2 := 0, len(nums)-1
	zeroNum := 0
	for i := 0; i+zeroNum < len(nums); i++ {
		for i+zeroNum < len(nums) && nums[i+zeroNum] == 0 {
			zeroNum++
		}
		if i+zeroNum < len(nums) {
			nums[i] = nums[i+zeroNum]
		}
	}
	for i := len(nums) - zeroNum; i < len(nums); i++ {
		nums[i] = 0
	}
}

func TestMoveZeros(t *testing.T) {
	data := []int{0, 1, 0, 3, 1, 2}
	moveZeroes(data)
	log.Println(data)
}
