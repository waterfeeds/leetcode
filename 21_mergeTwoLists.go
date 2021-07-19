package main

func main() {

}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	var preHead *ListNode = &ListNode{}
	var p *ListNode = preHead
	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			p.Next = l1
			l1 = l1.Next
			p = p.Next
		} else {
			p.Next = l2
			l2 = l2.Next
			p = p.Next
		}
	}

	if l1 != nil {
		p.Next = l1
	}
	if l2 != nil {
		p.Next = l2
	}

	return preHead.Next
}
