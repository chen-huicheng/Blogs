package leetcode

func permute(nums []int) [][]int {
	vis := make([]bool, len(nums))
	res := make([][]int, 0)
	path := make([]int, len(nums))
	var dfs func(int)
	dfs = func(i int) {
		if i == len(nums) {
			res = append(res, append([]int(nil), path...))
			return
		}
		for j, v := range nums {
			if vis[j] {
				continue
			}
			path[i] = v
			vis[j] = true
			dfs(i + 1)
			vis[j] = false
		}
	}

	dfs(0)
	return res
}
