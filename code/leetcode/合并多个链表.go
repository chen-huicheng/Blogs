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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var h ListNode
	n := &h
	p, q := list1, list2
	for p != nil && q != nil {
		if p.Val <= q.Val {
			n.Next = p
			p = p.Next
		} else {
			n.Next = q
			q = q.Next
		}
		n = n.Next
	}
	if p != nil {
		n.Next = p
	}
	if q != nil {
		n.Next = q
	}
	return h.Next
}
