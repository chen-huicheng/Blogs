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

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	res := make([][]int, 0)
	que := make([]*TreeNode, 0)
	que = append(que, root)
	l2r := true
	for len(que) > 0 {
		nextQue := make([]*TreeNode, 0)
		re := make([]int, len(que))
		for i, v := range que {
			if l2r {
				re[i] = v.Val
			} else {
				re[len(re)-1-i] = v.Val
			}
			nextQue = myAppend(nextQue, v.Left)
			nextQue = myAppend(nextQue, v.Right)
		}
		res = append(res, re)
		que = nextQue
		l2r = !l2r
	}
	return res

}
func myAppend(arr []*TreeNode, v *TreeNode) []*TreeNode {
	if v == nil {
		return arr
	}
	return append(arr, v)
}

//    1
//  2   3
// 4 5 6 7

// [1]
// [3,3]
// [7,6,5,4]
