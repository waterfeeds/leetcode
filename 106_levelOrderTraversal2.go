package main

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var result = make([][]int, 0)
	var queen = []*TreeNode{root}
	for len(queen) > 0 {
		var lenq = len(queen)
		solution := []int{}
		for lenq > 0 {
			current := queen[0]
			queen = queen[1:]
			solution = append(solution, current.Val)
			if current.Left != nil {
				queen = append(queen, current.Left)
			}
			if current.Right != nil {
				queen = append(queen, current.Right)
			}
			lenq--
		}
		result = append([][]int{solution}, result...)
	}

	return result
}
