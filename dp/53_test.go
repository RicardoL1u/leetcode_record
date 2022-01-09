package dp

func maxSubArray(nums []int) int {
	pre, ans := 0, nums[0]
	for _, v := range nums {
		// 前半部是个累赘 就直接丢掉
		// 给定长度为 k 的arr
		// pre = f[k]
		// f[k+1] = max(nums[k+1],f[k]+nums[k+1])
		pre = max(v, pre+v)
		ans = max(pre, ans)
	}
	return ans
}

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	} else {
// 		return b
// 	}
// }
