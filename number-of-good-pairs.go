// https://leetcode.com/problems/number-of-good-pairs

func numIdenticalPairs(nums []int) int {
	numMap := make(map[int]int)

	for _, val := range nums {
		numMap[val]++
	}

	var result int

	for _, val := range numMap {
		if val > 1 {
			result += (val) * (val - 1) / 2
		}
	}

	return result
}
