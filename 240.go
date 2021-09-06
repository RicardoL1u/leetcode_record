package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	sr, sc := 0, 0
	er := len(matrix)
	ec := len(matrix[0])

	for sr < er && sc < ec {
		mr := (er-sr)/2 + sr
		mc := (ec-sc)/2 + sc
		// fmt.Println(sr, sc, er, ec, mr, mc)
		if matrix[mr][mc] == target {
			return true
		} else if matrix[mr][mc] < target {
			col := copyCol(matrix, mc)
			m1 := binarySearch(col, target, mc, false)
			m2 := binarySearch(matrix[mr], target, mr, false)
			// fmt.Println(m1, m2)
			if m1 || m2 {
				return true
			}
			sr = mr + 1
			sc = mc + 1
		} else {
			col := copyCol(matrix, mc)
			m1 := binarySearch(col, target, mc, true)
			m2 := binarySearch(matrix[mr], target, mr, true)
			// fmt.Println(m1, m2)
			if m1 || m2 {
				return true
			}
			er = mr
			ec = mc
		}
	}

	return false
}

// var cnt int = 0

func binarySearch(nums []int, target, p int, isBig bool) bool {
	// cnt := 0
	s := 0
	e := len(nums)
	if isBig {
		e = p
	} else {
		s = p + 1
	}
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
	mat := [][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}
	fmt.Println(mat)
	fmt.Println(mat[1])
	// fmt.Println(mat[][1])
	fmt.Println(copyCol(mat, 1))
	// fmt.Println(searchMatrix(mat, 4))
	fmt.Println(searchMatrix(mat, 4))
}
