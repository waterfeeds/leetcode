package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func main() {

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeList := make([]*ListNode, 0)
	var p = head
	for p != nil {
		nodeList = append(nodeList, p)
		p = p.Next
	}

	if n == 1 {
		if len(nodeList) == 1 {
			return nil
		}
		nodeList[len(nodeList)-2].Next = nil
	} else if n == len(nodeList) {
		head = nodeList[1]
	} else {
		nodeList[len(nodeList)-n-1].Next = nodeList[len(nodeList)-n+1]
	}

	return head
}
