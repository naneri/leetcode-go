// https://leetcode.com/problems/monotonic-array

func isMonotonic(nums []int) bool {
	if len(nums) == 1 {
		return true
	}
	var direction int

	for i := 1; i < len(nums); i++ {
		if nums[i] < nums[i-1] {
			if direction == 1 {
				return false
			}
			if direction == 0 {
				direction = -1
			}
		}
		
		if nums[i] > nums[i-1] {
			if direction == -1 {
				return false
			}
			if direction == 0 {
				direction = 1
			}
		}
	}
	return true
}
