// https://leetcode.com/problems/validate-binary-tree-nodes

type TreeNodeN struct {
	Val       int
	HasParent bool
	Left      *TreeNodeN
	Right     *TreeNodeN
}

func validateBinaryTreeNodes(n int, leftChild []int, rightChild []int) bool {
	nodes := make([]TreeNodeN, len(leftChild), len(leftChild))
	notIsHead := make([]bool, len(leftChild), len(leftChild))

	for i := 0; i < len(leftChild); i++ {
	if leftChild[i] != -1 {
			if nodes[leftChild[i]].HasParent {
				return false
			}
			nodes[i].Left = &nodes[leftChild[i]]
			nodes[leftChild[i]].HasParent = true
			notIsHead[leftChild[i]] = true
		
		}

		if rightChild[i] != -1 {

			if nodes[rightChild[i]].HasParent {
				return false
			}
			nodes[i].Right = &nodes[rightChild[i]]
			nodes[rightChild[i]].HasParent = true
			notIsHead[rightChild[i]] = true

		}
	}

	var head *TreeNodeN

	for key, node := range notIsHead {
		if node == false {
			if head != nil {
				return false
			}
			head = &nodes[key]
		}
	}

	if head == nil {
		return false
	}

	var countChildren func(node *TreeNodeN) int
	
	countChildren = func(node *TreeNodeN) int {
		if node == nil {
			return 0
		}
		
		return 1 + countChildren(node.Left) + countChildren(node.Right)
	}
	
	return countChildren(head) == len(leftChild)
}
