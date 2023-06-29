package leetcode

import (
	"math"
	"sort"
)

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 1
	for ; i > 0; i-- {
		if nums[i-1] < nums[i] {
			break
		}
	}
	if i == 0 {
		sort.Ints(nums)
		return
	}
	maxN := math.MaxInt
	j := i
	swapIdx := j
	for ; j < n; j++ {
		if nums[j] > nums[i-1] && nums[j] < maxN {
			maxN = nums[j]
			swapIdx = j
		}
	}
	nums[i-1], nums[swapIdx] = nums[swapIdx], nums[i-1]
	sort.Ints(nums[i:])
}
