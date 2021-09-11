package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	// sr, sc := 0, 0
	er := len(matrix)
	ec := len(matrix[0])
	for cnt := 0; cnt < er && cnt < ec; cnt++ {
		m1 := binarySearch(matrix[cnt], cnt, target)
		m2 := binarySearch(copyCol(matrix, cnt), cnt, target)
		if m1 || m2 {
			return true
		}
	}
	return false
}

// var cnt int = 0

func binarySearch(nums []int, s, target int) bool {
	// cnt := 0
	// s := 0
	e := len(nums)

	for s < e {
		m := (e-s)/2 + s
		if nums[m] < target {
			s = m + 1
		} else {
			e = m
		}
	}
	if s >= len(nums) {
		return false
	} else {
		return nums[s] == target
	}
}

func copyCol(mat [][]int, col int) []int {
	res := []int{}
	// mxn
	m := len(mat)
	for i := 0; i < m; i++ {
		res = append(res, mat[i][col])
	}
	return res
}

// func binarySearch()

func main() {
	// mat := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	// mat := [][]int{{1, 4, 7, 11, 15}}
	mat := [][]int{{3, 3, 8, 13, 13, 18}, {4, 5, 11, 13, 18, 20}, {9, 9, 14, 15, 23, 23}}
	// mat := [][]int{{1}, {2}, {3}, {10}, {18}}

	fmt.Println(copyCol(mat, 0))
	for len := range mat {
		fmt.Println(mat[len])
	}
	// fmt.Println(searchMatrix(mat, 4))
	fmt.Println(searchMatrix(mat, 27))
}
