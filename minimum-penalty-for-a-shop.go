// https://leetcode.com/problems/minimum-penalty-for-a-shop


func bestClosingTime(customers string) int {
	customersSlice := []rune(customers)
	maxPenalty := 0
	res := len(customers)
	var totalScore int

	for i := len(customersSlice) - 1; i >= 0; i-- {
		if customersSlice[i] == 'Y' {
			totalScore++
		} else {
			totalScore--
		}

		if totalScore <= maxPenalty {
			maxPenalty = totalScore
			res = i
		}
	}

	return res
}
