package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lA := lenList(headA)
	lB := lenList(headB)
	if lA < lB {
		headA, headB = headB, headA
		lA, lB = lB, lA
	}
	for i := 0; i < lA-lB; i++ {
		headA = headA.Next
	}
	for headA != nil && headB != nil {
		if headB == headA {
			return headA
		}
		headB = headB.Next
		headA = headA.Next
	}
	return nil
}

func lenList(head *ListNode) int {
	res := 0
	for head != nil {
		res++
		head = head.Next
	}
	return res
}
