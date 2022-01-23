package dfs

import (
	"log"
	"testing"
)

func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	var dfs func(uint8, []int)
	dfs = func(used uint8, now []int) {
		if used == 1<<len(nums)-1 {
			tmp := make([]int, len(nums))
			copy(tmp, now)
			ans = append(ans, tmp)
			// log.Println(used, now, ans)
			// return ans
		}

		for i := 0; i < len(nums); i++ {
			if ((1 << i) & used) == 0 {
				now = append(now, nums[i])
				dfs(used|(1<<i), now)
				now = now[:len(now)-1]
			}
		}
	}
	dfs(0, make([]int, 0))
	return ans
}

// func dfs(used uint8, nums, now []int) [][]int {
// 	ans := make([][]int, 0)
// 	if used == 1<<len(nums)-1 {
// 		log.Println(used, now)
// 		ans = append(ans, now)
// 		return ans
// 	}

// 	for i := 0; i < len(nums); i++ {
// 		if ((1 << i) & used) == 0 {
// 			now = append(now, nums[i])
// 			ans = append(ans, dfs(used|(1<<i), nums, now)...)
// 			now = now[:len(now)-1]
// 		}
// 	}
// 	return ans
// }

// var cnt int = 0

// func dfs(used []bool, nums, now []int) [][]int {
// 	ans := make([][]int, 0)
// 	// log.Println(used, now, cnt)
// 	// cpy := make([]bool, len(nums))
// 	mark := false
// 	for k, v := range used {
// 		if !v {
// 			// v = true
// 			cpy := make([]bool, len(used))
// 			copy(cpy, used)
// 			cpy[k] = true
// 			// log.Println(cpy, used)
// 			mark = true
// 			// now = append(now, nums[k])
// 			tmp := make([]int, len(now))
// 			copy(tmp, now)
// 			tmp = append(tmp, nums[k])
// 			ans = append(ans, dfs(cpy, nums, tmp)...)
// 		}
// 	}
// 	if !mark {
// 		// log.Printf("this is the end part with now =%v", now)
// 		ans = append(ans, now)
// 		return ans
// 	}
// 	return ans
// }

func TestPer(t *testing.T) {
	log.Println(permute([]int{1, 2, 3, 4}))
}
