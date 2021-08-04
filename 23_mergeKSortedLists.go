package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
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

func main() {
	l1 := []int{-8,-7,-6,-3,-2,-2,0,3}
	l2 := []int{-10,-6,-4,-4,-4,-2,-1,4}
	l3 := []int{-10,-9,-8,-8,-6}
	l4 := []int{-10,0,4}
	lists := make([]*ListNode, 0)
	lists = append(lists, makeNodeList(l1), makeNodeList(l2), makeNodeList(l3), makeNodeList(l4))
	fmt.Println(mergeKLists(lists))
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) < 1 {
		return nil
	}

	return recursionKList(lists, 0, len(lists)-1)
}

func recursionKList(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}

	mid := (left+right)/2
	leftList := recursionKList(lists, left, mid)
	rightList := recursionKList(lists, mid+1, right)
	return mergeTwoLists(leftList, rightList)
}
