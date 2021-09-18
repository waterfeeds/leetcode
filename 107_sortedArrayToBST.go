package main

func sortedArrayToBST(nums []int) *TreeNode {
	return sortedRecursion(nums, 0, len(nums)-1)
}

func sortedRecursion(nums []int, left int, right int) *TreeNode {
	if left > right {
		return nil
	}

	var mid = (left+right+1)/2
	var root = &TreeNode{Val: nums[mid]}
	root.Left = sortedRecursion(nums, left, mid-1)
	root.Right = sortedRecursion(nums, mid+1, right)

	return root
}
