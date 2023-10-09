package leetcode

/*

 */
func firstMissingPositive(nums []int) int {
	n := len(nums)
	for i, v := range nums {
		if v <= 0 || v > n || v == i+1 {
			continue
		}
		nums[i] = 0
		k := v
		for k > 0 && k <= n && nums[k-1] != k {
			nums[k-1], k = k, nums[k-1]
		}
	}
	for i, v := range nums {
		if i+1 != v {
			return i + 1
		}
	}
	return n + 1
}
