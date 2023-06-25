package leetcode

func trap(height []int) int {
	n := len(height)
	lMax := make([]int, len(height))
	rMax := make([]int, len(height))
	lm, rm := 0, 0
	for i := range height {
		lm = max(lm, height[i])
		rm = max(rm, height[n-i-1])
		lMax[i] = lm
		rMax[n-i-1] = rm
	}
	res := 0
	for i := 0; i < n; i++ {
		res += min(lMax[i], rMax[i]) - height[i]
	}
	return res
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
