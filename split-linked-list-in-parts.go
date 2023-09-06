// https://leetcode.com/problems/split-linked-list-in-parts

func splitListToParts(head *ListNode, k int) []*ListNode {
	var listSlice []*ListNode
	for head != nil {
		listSlice = append(listSlice, head)
		head = head.Next
	}

	sliceSize := len(listSlice) / k
	leftOver := len(listSlice) % k

	var res []*ListNode

	var index int
	for index < len(listSlice) {
		res = append(res, listSlice[index])
		index = index + sliceSize
		if leftOver > 0 {
			index++
			leftOver--
		}
		listSlice[index - 1].Next = nil
	}

	for len(res) < k {
		res = append(res, nil)
	}
	return res
}
