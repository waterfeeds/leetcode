package main

type ListNode struct {
  Val int
  Next *ListNode
}

func main() {
	l1 := []int{5}
	l2 := []int{5}

	addTwoNumbers(makeNodeList(l1), makeNodeList(l2))
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var p *ListNode
	var highVal = 0

	for l1 != nil && l2 != nil {
		sumVal := l1.Val+l2.Val+highVal
		lowVal := sumVal%10
		highVal = sumVal/10

		if head == nil {
			head = &ListNode{
				Val:  lowVal,
				Next: nil,
			}
			p = head
			l1 = l1.Next
			l2 = l2.Next
		} else {
			p.Next = &ListNode{
				Val: lowVal,
			}
			p = p.Next
			l1 = l1.Next
			l2 = l2.Next
		}
	}

	for l1 != nil {
		sumVal := l1.Val+highVal
		lowVal := sumVal%10
		highVal = sumVal/10
		p.Next = &ListNode{Val: lowVal}
		p = p.Next
		l1 = l1.Next
	}

	for l2 != nil {
		sumVal := l2.Val+highVal
		lowVal := sumVal%10
		highVal = sumVal/10
		p.Next = &ListNode{Val: lowVal}
		p = p.Next
		l2 = l2.Next
	}

	if highVal > 0 {
		p.Next = &ListNode{Val: highVal}
	}

	return head
}

func makeNodeList(nums []int) *ListNode {
	var head *ListNode
	var p *ListNode

	head = &ListNode{
		Val: nums[0],
	}
	p = head
	if len(nums) <= 1 {
		return head
	}
	for idx, num := range nums {
		if idx == 0 {
			continue
		}

		p.Next = &ListNode{
			Val: num,
		}
		p = p.Next
	}

	return head
}