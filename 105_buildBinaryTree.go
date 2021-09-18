package main

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	var root = &TreeNode{Val: preorder[0]}
	var stack = []*TreeNode{root}
	var index = 0
	for i := 1; i < len(preorder); i++ {
		if stack[len(stack)-1].Val != inorder[index] {
			leftNode := &TreeNode{Val: preorder[i]}
			stack[len(stack)-1].Left = leftNode
			stack = append(stack, leftNode)
		} else {
			var node *TreeNode
			for len(stack) > 0 && inorder[index] == stack[len(stack)-1].Val {
				index++
				node = stack[len(stack)-1]
				stack = stack[:len(stack)-1]
			}
			rightNode := &TreeNode{Val: preorder[i]}
			node.Right = rightNode
			stack = append(stack, rightNode)
		}
	}

	return root
}
