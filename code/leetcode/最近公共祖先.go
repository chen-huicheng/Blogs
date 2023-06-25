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

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	_, _, t := findTreeNode(root, p, q)
	return t
}

func findTreeNode(root, p, q *TreeNode) (bool, bool, *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	la, lb, lt := findTreeNode(root.Left, p, q)
	if la && lb {
		return la, lb, lt
	}
	ra, rb, rt := findTreeNode(root.Right, p, q)
	if ra && rb {
		return ra, rb, rt
	}
	a, b := la || ra, lb || rb
	if root == p {
		a = true
	}
	if root == q {
		b = true
	}
	if a && b {
		return a, b, root
	}
	return a, b, nil
}
