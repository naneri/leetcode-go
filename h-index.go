// https://leetcode.com/problems/h-index/

func hIndex(citations []int) int {
	sort.Slice(citations, func(i, j int) bool {
		return citations[i] < citations[j]
	})

	var index int

	for i := len(citations) - 1; i >= 0; i-- {
		if citations[i] >= len(citations)-i {
			index = len(citations) - i
		}
		
	}

	return index
}
