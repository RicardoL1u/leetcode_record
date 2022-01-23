package dp

func maxProfit(prices []int) int {
	minn := 1 << 30
	ans := -minn
	for _, v := range prices {
		// minn = min(v, minn)
		if v < minn {
			minn = v
		}
		ans = max(ans, v-minn)

	}
	return ans
}
