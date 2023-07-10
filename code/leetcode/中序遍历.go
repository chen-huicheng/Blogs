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

func inorderTraversal(root *TreeNode) []int {
	res := []int{}
	var trav func(*TreeNode)
	trav = func(root *TreeNode) {
		if root == nil {
			return
		}
		trav(root.Left)
		res = append(res, root.Val)
		trav(root.Right)
	}
	trav(root)
	return res
}

type point struct {
	root *TreeNode
	flag int
}

func inorderTraversal_1(root *TreeNode) []int {
	res := []int{}
	sta := []point{}
	sta = append(sta, point{root: root, flag: 0})
	for len(sta) > 0 {
		n := len(sta)
		p := sta[n-1]
		sta = sta[:n-1]
		if p.flag == 0 {
			if p.root.Left != nil {
				sta = append(sta, point{root: root, flag: 0})
			}
		}
	}
	trav(root)
	return res
}

func pushBack(a []point) b {

}
