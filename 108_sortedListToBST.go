package main

func sortedListToBST(head *ListNode) *TreeNode {
	arr := make([]int, 0)

	for head != nil {
		arr = append(arr, head.Val)
		head = head.Next
	}

	return sortedRecursion(arr, 0, len(arr)-1)
}
