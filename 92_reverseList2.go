package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	var preHead = &ListNode{Next: head}
	var preReverse = &ListNode{}
	var current = 1
	var p = preHead
	var q = head
	var tail *ListNode

	for q != nil {
		if current > right {
			break
		}
		if current >= left && current <= right {
			if tail == nil {
				tail = q
			}
			var temp = q.Next
			q.Next = preReverse.Next
			preReverse.Next = q
			q = temp
			current++
			continue
		}
		if current < left {
			p = p.Next
			q = q.Next
			current++
			continue
		}
		break
	}

	p.Next = preReverse.Next
	tail.Next = q
	return preHead.Next
}
