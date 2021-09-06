package main

import (
	"fmt"
)

var cnt int = 0
func binarySearch(s,e,x int, nums []int) int{
	m := (e + s) / 2
	fmt.Println(s," ",e," ",m," ",nums[m]," ")
	cnt ++
	if cnt == 15 {
		return 0
	}
	if e <= s {
		return s
	}
	if nums[m] < x {
		return binarySearch(m+1,e,x,nums);
	} else if nums[m] > x {
		return binarySearch(s,m,x,nums)
	} else {
		return binarySearch(s,m,x,nums)
	}
	return 0

}

func main(){
	nums := []int{1,2,3,6,8,9,10,13,15}
	fmt.Println(binarySearch(0,len(nums),11,nums))
	// nums := []int{1,6,8,9,10}
	// fmt.Println(binarySearch(0,len(nums),9,nums))
	
}