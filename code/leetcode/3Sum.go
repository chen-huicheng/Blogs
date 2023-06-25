package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	res := make([][]int, 0)
	re := make([]int, 3)
	for i := range nums {
		if i != 0 && nums[i] == nums[i-1] {
			continue
		}
		re[0] = nums[i]
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[i]+nums[l]+nums[r] == 0 {
				re[1], re[2] = nums[l], nums[r]
				tmp := make([]int, 0)
				tmp = append(tmp, re...)
				res = append(res, tmp)
				n := nums[l]
				for nums[l] == n && l < r {
					l++
				}
				n = nums[r]
				for nums[r] == n && l < r {
					r--
				}
			}
			for l < r && nums[i]+nums[l]+nums[r] > 0 {
				r--
			}
			for l < r && nums[i]+nums[l]+nums[r] < 0 {
				l++
			}
		}
	}
	return res
}
