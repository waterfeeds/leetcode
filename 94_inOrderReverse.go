package main

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	var result = make([]int, 0)
	inorderDfs(root, &result)
	return result
}

func inorderDfs(node *TreeNode, result *[]int) {
	if node == nil {
		return
	}

	if node.Left != nil {
		inorderDfs(node.Left, result)
	}

	*result = append(*result, node.Val)

	if node.Right != nil {
		inorderDfs(node.Right, result)
	}
}
