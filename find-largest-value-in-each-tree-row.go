// https://leetcode.com/problems/find-largest-value-in-each-tree-row/

func largestValues(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	return bfs([]*TreeNode{root})
}

func bfs(nodes []*TreeNode) []int {
	nextLevelNodes := make([]*TreeNode, 0)
	localVals := make([]int, 0)

	for _, val := range nodes {
			localVals = append(localVals, val.Val)
			if val.Left != nil {
				nextLevelNodes = append(nextLevelNodes, val.Left)
			}
			if val.Right != nil {
				nextLevelNodes = append(nextLevelNodes, val.Right)
			}
	}

	localMaxVal := maxInSlice(localVals)

	result := []int{localMaxVal}
	if len(nextLevelNodes) == 0 {
		return result
	}
	return append(result, bfs(nextLevelNodes)...)
}

func maxInSlice(ints []int) int {
	maxVal := ints[0]

	for _, val := range ints {
		if val > maxVal {
			maxVal = val
		}
	}

	return maxVal
}
