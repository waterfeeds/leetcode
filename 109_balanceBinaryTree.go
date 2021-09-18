package main

func isBalanced(root *TreeNode) bool {
	return height(root) >= 0
}

func height(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var left, right int

	left = height(root.Left)
	if left < 0 {
		return -1
	}
	right = height(root.Right)
	if right < 0 {
		return -1
	}

	if diff(left, right) > 1 {
		return -1
	}

	return max(left, right) + 1
}

func diff(a, b int) int {
	if a >= b {
		return a-b
	}

	return b-a
}
