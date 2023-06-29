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

func rightSideView(root *TreeNode) []int {
	res := []int{}
	if root == nil {
		return res
	}
	que := []*TreeNode{root}
	for len(que) != 0 {
		nextQ := []*TreeNode{}
		for i, p := range que {
			if i == len(que)-1 {
				res = append(res, p.Val)
			}
			if p.Left != nil {
				nextQ = append(nextQ, p.Left)
			}
			if p.Right != nil {
				nextQ = append(nextQ, p.Right)
			}
		}
		que = nextQ
	}
	return res
}
