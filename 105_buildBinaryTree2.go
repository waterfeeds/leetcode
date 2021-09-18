package main

func buildTree2(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}

	var root = &TreeNode{Val: postorder[len(postorder)-1]}
	var stack = []*TreeNode{root}
	var index = len(inorder)-1
	for i := len(postorder)-2; i >=0; i-- {
		if stack[len(stack)-1].Val != inorder[index] {
			rightNode := &TreeNode{Val: postorder[i]}
			stack[len(stack)-1].Right = rightNode
			stack = append(stack, rightNode)
		} else {
			var node *TreeNode
			for len(stack) > 0 && stack[len(stack)-1].Val == inorder[index] {
				node = stack[len(stack)-1]
				index--
				stack = stack[:len(stack)-1]
			}
			leftNode := &TreeNode{Val: postorder[i]}
			node.Left = leftNode
			stack = append(stack, leftNode)
		}
	}

	return root
}
