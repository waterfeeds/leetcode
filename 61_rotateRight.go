package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if k == 0 || head == nil {
		return head
	}


	var fast, slow = head, head
	var length = 1

	for fast.Next != nil {
		if length > k {
			slow = slow.Next
		}
		fast = fast.Next
		length++
	}

	if length == k {
		return head
	}

	if length > k {
		fast.Next = head
		newHead := slow.Next
		slow.Next = nil
		return newHead
	}

	var rotate = k % length
	if rotate == 0 {
		return head
	}

	var rotateIndex = 1
	for rotateIndex < length-rotate {
		slow = slow.Next
		rotateIndex++
	}

	fast.Next = head
	newHead := slow.Next
	slow.Next = nil
	return newHead
}
