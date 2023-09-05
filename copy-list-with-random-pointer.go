// https://leetcode.com/problems/copy-list-with-random-pointer

func copyRandomList(head *Node) *Node {

	if head == nil {
		return nil
	}
	nodes := make(map[*Node]*Node)

	newHead := &Node{Val: head.Val}
	headCopy := head
	newHeadCopy := newHead
	nodes[head] = newHead
	
	for {
		head = head.Next
		if head != nil {
			newNode := &Node{Val: head.Val}
			newHead.Next = newNode
			newHead = newNode
			nodes[head] = newHead
		} else {
			break
		}
	}
	
	res := newHeadCopy
	for {
		if headCopy == nil {
			break
		}
		if headCopy.Random != nil {
			newHeadCopy.Random = nodes[headCopy.Random]
		}	
		headCopy = headCopy.Next
		newHeadCopy = newHeadCopy.Next
	}
	
	return res
}
