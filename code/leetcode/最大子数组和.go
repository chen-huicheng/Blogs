package leetcode

func maxSubArray(nums []int) int {
	maxN := nums[0]
	sum := 0
	for _, n := range nums {
		sum += n
		if sum > maxN {
			maxN = sum
		}
		if sum < 0 {
			sum = 0
		}
	}
	return maxN
}
