package main

func partition(head *ListNode, x int) *ListNode {
	var bigPreHead = &ListNode{}
	var smallPreHead = &ListNode{}

	var p = head
	var pBig = bigPreHead
	var pSmall = smallPreHead
	for p != nil {
		if p.Val >= x {
			pBig.Next = p
			pBig = pBig.Next
			p = p.Next
		} else {
			pSmall.Next = p
			pSmall = pSmall.Next
			p = p.Next
		}
	}

	pBig.Next = nil
	pSmall.Next = bigPreHead.Next
	return smallPreHead.Next
}
