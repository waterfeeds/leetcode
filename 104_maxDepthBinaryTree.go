package main

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	var maxL = 0
	var queen = []*TreeNode{root}
	for len(queen) > 0 {
		var lenq = len(queen)
		for lenq > 0 {
			current := queen[0]
			queen = queen[1:]
			if current.Left != nil {
				queen = append(queen, current.Left)
			}
			if current.Right != nil {
				queen = append(queen, current.Right)
			}
			lenq--
		}
		maxL++
	}

	return maxL
}
