// https://leetcode.com/problems/third-maximum-number/



func thirdMax(nums []int) int {
	intMap := make(map[int]bool)
	newNums := make([]int, 0)

	for _, val := range nums {
		if _, ok := intMap[val]; !ok {
			intMap[val] = true
			newNums = append(newNums, val)
		}
	}

	sort.Ints(newNums)
	if len(newNums) < 3 {
		return newNums[len(newNums)-1]
	}

	return newNums[len(newNums) - 3]
}
