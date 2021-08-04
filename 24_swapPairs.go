package main

func main() {

}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var prevHead = &ListNode{}
	prevHead.Next = head.Next
	var prev, p, q = prevHead, head, head.Next
	for p != nil && q != nil {
		p.Next = q.Next
		q.Next = p
		prev.Next = q

		prev = p
		p = p.Next
		if p == nil || p.Next == nil {
			break
		}
		q = p.Next
	}

	return prevHead.Next
}
