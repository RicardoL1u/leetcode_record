package main

import (
	"fmt"
)

func search(nums []int, target int) int {
	oriNums := make([]int,5010)
	n := len(nums)
	if n == 1{
		if nums[0] != target{
			return -1
		} else {
			return 0
		}
	}
	k := findRotatedPoi(nums)
	copy(oriNums[0:n-k],nums[k:])
	copy(oriNums[n-k:],nums[0:k])
	// k := 10
	// fmt.Println(oriNums[0:n])
	oriPos := findTarget(oriNums,target,n)
	if oriPos == -1 {
		return oriPos
	}
	nowPos := getOriPos(oriPos,n,k)
	// fmt.Println(oriPos, " ", nowPos)
	return nowPos
}

func getOriPos (oriPos,n,k int) int{
	res := 0
	if oriPos < n - k {
		res = oriPos + k
	} else {
		res = oriPos - (n - k)
	}
	return res
}

func findTarget(nums[] int,target int,l int) int {
	s := 0
	e := l
	cmpVal := target
	for s < e {
		p := (e - s) / 2 + s
		// fmt.Println(s," ",e," ",p)
		if cmpVal > nums[p]{
			s = p + 1
		} else {
			e = p
		}
	}
	if nums[s] == cmpVal {
		return s
	} else {
		return -1
	}
}

func findRotatedPoi(nums[] int) int {
	s := 0
	e := len(nums)
	cmpVal := nums[e - 1]
	for s < e {
		p := (e - s) / 2 + s
		if cmpVal < nums[p]{
			s = p + 1
		} else {
			e = p
		}
	}
	return s
}

func main(){

	// nums := []int{4,5,6,7,0,1,2}
	// nums := []int{11,22,33,44,55,3,5,7}
	// nums := []int{1}
	a := []int{4,5,6,7,0,1,2}
	// b := make([]int,10)
	// cp := copy(b[4:6],a[1:3])
	// fmt.Println(cp)
	// fmt.Println(b)
	// fmt.Println(findRotatedPoi(nums))
	fmt.Println(search(a,1))
}