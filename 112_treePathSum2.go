package main

func pathSum(root *TreeNode, targetSum int) [][]int {
	var result = make([][]int, 0)
	pathSumDFS(root, targetSum, []int{}, &result)
	return result
}

func pathSumDFS(root *TreeNode, targetSum int, solution []int, result *[][]int) {
	if root == nil {
		return
	}

	if targetSum == root.Val && root.Left == nil && root.Right == nil {
		*result = append(*result, append(copylist(solution), root.Val))
		return
	}

	if root.Left != nil {
		pathSumDFS(root.Left, targetSum-root.Val, append(solution, root.Val), result)
	}
	if root.Right != nil {
		pathSumDFS(root.Right, targetSum-root.Val, append(solution, root.Val), result)
	}
}
