package leetcode

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

func levelOrder(root *TreeNode) [][]int {
	res := [][]int{}
	if root == nil {
		return res
	}
	que := []*TreeNode{root}
	for len(que) != 0 {
		level := []int{}
		nextQ := []*TreeNode{}
		for _, p := range que {
			level = append(level, p.Val)
			if p.Left != nil {
				nextQ = append(nextQ, p.Left)
			}
			if p.Right != nil {
				nextQ = append(nextQ, p.Right)
			}
		}
		que = nextQ
		res = append(res, level)
	}
	return res
}
