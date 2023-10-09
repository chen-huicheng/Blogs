package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var p, lp, rp *ListNode
	var h ListNode
	h.Next = head
	i := 0
	p = &h
	for p != nil {
		if i+1 == left {
			lp = p
		}
		if i == right {
			rp = p
			break
		}
		p = p.Next
		i++
	}
	p = rp.Next
	rp.Next = nil
	rh := lp.Next
	lp.Next = reverse(lp.Next)
	rh.Next = p
	return h.Next
}

func reverse(h *ListNode) *ListNode {
	var res, p *ListNode
	for h != nil {
		p = h
		h = h.Next
		p.Next = res
		res = p
	}
	return res
}
