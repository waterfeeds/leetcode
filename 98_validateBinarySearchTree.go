package main

func isValidBST(root *TreeNode) bool {
	var valid = true
	isValidBSTDfs(root, 1<<32-1, -1<<32, &valid)
	return valid
}

func isValidBSTDfs(node *TreeNode, maxV int, minV int, valid *bool) {
	if !*valid {
		return
	}

	if node == nil {
		return
	}

	if node.Val >= maxV || node.Val <= minV {
		*valid = false
		return
	}

	if node.Left != nil {
		isValidBSTDfs(node.Left, min(maxV, node.Val), minV, valid)
	}

	if node.Right != nil {
		isValidBSTDfs(node.Right, maxV, max(minV, node.Val), valid)
	}
}
