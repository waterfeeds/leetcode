package main

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var minL = 0
	var queen = []*TreeNode{root}
	for len(queen) > 0 {
		var lenq = len(queen)
		for lenq > 0 {
			current := queen[0]
			queen = queen[1:]
			if current.Left == nil && current.Right == nil {
				return minL+1
			}
			if current.Left != nil {
				queen = append(queen, current.Left)
			}
			if current.Right != nil {
				queen = append(queen, current.Right)
			}
			lenq--
		}
		minL++
	}

	return minL
}
