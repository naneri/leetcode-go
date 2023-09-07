// https://leetcode.com/problems/reverse-linked-list-ii

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	index := 1
	newHead := head
	var startReverse bool
	var values 	[]int

	var leftNode, rightNode *ListNode

	for leftNode == nil || rightNode == nil {
		if head == nil || index > right {
			break
		}

		if index == left {
			startReverse = true
		}

		if startReverse {
			values = append(values, head.Val)
		}

		index++
		head = head.Next
	}

	newIndex := 1
	resHead := newHead

	for newHead != nil {
		if newIndex >= left && newIndex <= right {
			newHead.Val = values[len(values) - 1 - newIndex + left]
		}

		newIndex++
		newHead = newHead.Next
	}

	return resHead
}
