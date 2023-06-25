package leetcode

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, v := range nums {
		m[v] = i
	}
	for i, v := range nums {
		if v, ok := m[target-v]; ok {
			return []int{v, i}
		}
	}
	return nil
}
