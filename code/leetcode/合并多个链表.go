/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
package leetcode

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {

	getMinValPtr := func() *ListNode {
		minVal := math.MaxInt
		resIdx := -1
		for i, v := range lists {
			if v != nil && v.Val < minVal {
				minVal = v.Val
				resIdx = i
			}
		}
		if resIdx == -1 {
			return nil
		}
		v := lists[resIdx]
		lists[resIdx] = v.Next
		return v
	}
	head := ListNode{Next: nil}
	nextPtr := &head
	for {
		v := getMinValPtr()
		if v == nil {
			break
		}
		nextPtr.Next = v
		nextPtr = nextPtr.Next
	}
	nextPtr.Next = nil
	return head.Next
}
