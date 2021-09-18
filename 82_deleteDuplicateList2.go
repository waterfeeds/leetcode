package main

func deleteDuplicatesList2(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	var preHead = &ListNode{Next: head}
	var p, q = preHead, head
	var toRemove = -101
	for q != nil {
		if q.Val == toRemove {
			p.Next = q.Next
			q = p.Next
			continue
		}
		if q.Next == nil {
			break
		}

		if q.Val == q.Next.Val {
			toRemove = q.Val
			p.Next = q.Next
			q = p.Next
			continue
		}

		p = p.Next
		q = q.Next
	}

	return preHead.Next
}
