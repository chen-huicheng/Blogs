package leetcode

import "math"

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		lm := dfs(root.Left)
		rm := dfs(root.Right)
		maxSum = max(maxSum, lm+rm+root.Val)
		return max(0, max(lm, rm)+root.Val)
	}
	dfs(root)
	return maxSum
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
