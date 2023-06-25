package leetcode

func maxProfit(prices []int) int {
	minP := prices[0]
	res := 0
	for _, v := range prices {
		res = max(res, v-minP)
		minP = min(minP, v)
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
