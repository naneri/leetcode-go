// https://leetcode.com/problems/linked-list-cycle

func hasCycle(head *ListNode) bool {
	nodeMap := make(map[*ListNode]bool)
	
	for head != nil {
		if nodeMap[head] != false {
			return true
		}
		
		nodeMap[head] = true
		head = head.Next
	}
	
	return false
}
