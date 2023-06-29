package leetcode

func lengthOfLIS(nums []int) int {
	k := 0
	n := len(nums)
	for i := 1; i < n; i++ {
		if nums[i] > nums[k] {
			k++
			nums[k] = nums[i]
		}
		for j := k; j >= 0; j-- {
			if nums[i] > nums[j] {
				break
			}
			if j == 0 || (j > 0 && nums[j-1] < nums[i]) {
				nums[j] = nums[i]
			}
		}
	}
	return k + 1
}
