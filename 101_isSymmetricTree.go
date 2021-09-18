package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return isSymmetricDfs(root.Left, root.Right)
}

func isSymmetricDfs(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	return left.Val == right.Val && isSymmetricDfs(left.Left, right.Right) && isSymmetricDfs(left.Right, right.Left)
}
