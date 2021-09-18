package main

func generateTrees(n int) []*TreeNode {
	return backtraceGenetateTrees(1, n)
}

func backtraceGenetateTrees(start int, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}

	var trees = make([]*TreeNode, 0)
	for index := start; index <= end; index++ {
		leftTrees := backtraceGenetateTrees(start, index-1)
		rightTrees := backtraceGenetateTrees(index+1, end)
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				tree := &TreeNode{Val: index, Left: left, Right: right}
				trees = append(trees, tree)
			}
		}
	}

	return trees
}
