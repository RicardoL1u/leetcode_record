package leetcode

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	sr, sc := 0, 0
	er := len(matrix)
	ec := len(matrix[0])
	fmt.Println(er, " ", ec)
	fmt.Println(sr, " ", sc)

	for sr < er && sc < ec {
		mr := (er-sr)/2 + sr
		mc := (ec-sc)/2 + sc
		if matrix[mr][mc] == target {
			return true
		} else {

		}
	}

	return false
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
}
