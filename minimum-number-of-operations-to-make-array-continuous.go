
func minOperations(nums []int) int {
	numsLen := len(nums)

	nonDups := getNonDuplicatesSorted(nums)

	var right, currentCount, maxCount int

	for left := range nonDups {
		for right < len(nonDups) && nonDups[right] < nonDups[left]+numsLen {
			currentCount++
			right++
		}

		if currentCount > maxCount {
			maxCount = currentCount
		}

		currentCount--
	}

	return numsLen - maxCount
}

func getNonDuplicatesSorted(nums []int) []int {
	dups := make(map[int]int)

	for _, num := range nums {
		dups[num]++
	}

	nonDups := make([]int, 0)

	for key := range dups {
		nonDups = append(nonDups, key)
	}

	sort.Ints(nonDups)

	return nonDups
}
