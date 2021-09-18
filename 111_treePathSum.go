package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	return hasPathSumRecursion(root, targetSum)
}

func hasPathSumRecursion(root *TreeNode, current int) bool {
	if root == nil {
		return false
	}

	if root.Val == current && root.Left == nil && root.Right == nil {
		return true
	}

	return hasPathSumRecursion(root.Left, current-root.Val) || hasPathSumRecursion(root.Right, current-root.Val)
}
