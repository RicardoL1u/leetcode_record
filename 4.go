package main

import (
	"fmt"
)
// var req int
var m int
var n int

func unitMid(nums[] int) float64 {
	lenth := len(nums)
	if lenth % 2 == 1{
		return float64(nums[lenth/2])
	} else {
		return (float64(nums[lenth/2]) + float64(nums[lenth/2 - 1]))/2
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m = len(nums1)
	n = len(nums2)
	k := int((m+n)/2)
	if m == 0 {
		return unitMid(nums2)
	} else if n == 0 {
		return unitMid(nums1)
	}
	if (m+n)%2 == 1 {
		return binarySearch(0,0,m,n,k,nums1,nums2)
	} else {
		return (binarySearch(0,0,m,n,k-1,nums1,nums2) + binarySearch(0,0,m,n,k,nums1,nums2))/2
	}

}
var cnt int = 0
func binarySearch(s1,s2,e1,e2,k int,nums1,nums2 []int) float64{
	var p1 int = (e1 + s1) / 2
	var p2 int = (e2 + s2) / 2
	// cnt ++
	// if cnt == 15 {
	// 	return 0.0
	// }
	
	prev1Min := p1
	prev2Min := p2
	prev1Max := p1 + p2
	prev2Max := p1 + p2
	
	fmt.Println(s1," ",e1," ",s2," ",e2)
	if (s1 == e1 - 1 || s1 == e1) && (s2 == e2 - 1 || s2 == e2) {
		prev1 := p1
		prev2 := p2
		if nums1[p1] > nums2[p2] {
			prev1 += (prev2 + 1)
		} else if nums1[p1] < nums2[p2] {
			prev2 += (prev1 + 1)
		}
		// ans := 0
		if prev1 == k{
			return float64(nums1[p1])
		} else {
			return float64(nums2[p2])
		}
		// return p1,p2
	}
	if nums1[p1] > nums2[p2] {
		// prev1Max ++
		prev1Min = p1 + p2 + 1
		prev1Max = e2 + p1
	} else if nums1[p1] < nums2[p2]{
		// prev2Max ++	
		prev2Min = p1 + p2 + 1
		prev2Max = e1 + p2
	} else {
		prev1Min = p1 + p2
		prev2Min = p1 + p2
	}
	fmt.Println(s1," ",e1," ",s2," ",e2," ",prev1Max," ",prev1Min," ",prev2Max," ",prev2Min)
	if prev1Min == prev1Max && prev1Min == k {
		return float64(nums1[p1])
	}
	if prev2Min == prev2Max && prev2Min == k {
		return float64(nums2[p2])
	}
	if prev1Min > k{
		e1 = p1
	}
	if prev2Min > k {
		e2 = p2
	}
	if prev1Max <= k {
		s1 = p1
	}
	if prev2Max <= k {
		s2 = p2
	}
	return binarySearch(s1,s2,e1,e2,k,nums1,nums2);
	

}

func main(){
	nums1 := []int{1,5}
	nums2 := []int{2,3,4,6} 
	fmt.Println(findMedianSortedArrays(nums1,nums2))

}