// https://leetcode.com/problems/combination-sum-iii

func combinationSum3(k int, n int) [][]int {
	return makeCombo(k, n, 9)
}

func makeCombo(nums, total, maxNum int ) [][]int {
	if nums == 0 {
		return [][]int{}
	}

	if nums == 1 {
		if total <= maxNum && maxNum > 0 {
			return [][]int{{total}}
		}

		return [][]int{}
	}

	var result [][]int

	for i := maxNum; i > 0; i-- {
		var output [][]int
		if i <= total - 1 {
			output = makeCombo(nums - 1, total - i, i - 1)
		}

		for j := 0; j < len(output); j++ {
			output[j] = append(output[j], i)
		}

		result = append(result, output...)
	}

	return result
}
