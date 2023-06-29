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

func reorderList(head *ListNode) {
	n := 0
	p := head
	for p != nil {
		n++
		p = p.Next
	}
	p = head
	for i := 1; i < (n+1)/2; i++ {
		p = p.Next
	}
	h2 := p.Next
	p.Next = nil
	h2 = reverseList(h2)
	h1 := head
	res := ListNode{}
	p = &res
	for h1 != nil && h2 != nil {
		p.Next = h1
		p = p.Next
		h1 = h1.Next

		p.Next = h2
		p = p.Next
		h2 = h2.Next
	}
	p.Next = nil
	if h1 != nil {
		p.Next = h1
	}
	head = res.Next
}
func reverseList(head *ListNode) *ListNode {
	p := head
	var res *ListNode
	for p != nil {
		tmp := p.Next
		p.Next = res
		res = p
		p = tmp
	}
	return res
}
