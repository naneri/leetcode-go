// https://leetcode.com/problems/remove-duplicates-from-sorted-array-ii


func removeDuplicates(nums []int) int {
	var count int
	var dupCount int

	for key, val := range nums {
		if key == 0 {
			nums[count] = val
			count++
			continue
		}

		if val == nums[key - 1] {
			dupCount++
			if dupCount > 1{
				count--
			}
		} else {
			dupCount = 0
		}

		nums[count] = val
		count++
	}

	return count
}
