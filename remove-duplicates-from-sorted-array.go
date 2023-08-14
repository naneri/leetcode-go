// https://leetcode.com/problems/remove-duplicates-from-sorted-array


func removeDuplicates(nums []int) int {
	var count int
	resArr := make([]int, len(nums))
	for key, val := range nums {
		resArr[count] = val
		count++
		if key != 0 {
			if nums[key] == nums[key - 1] {
				count--
			} 
		}
	}
	
	for key, val := range resArr {
		nums[key] = val
	}
	
	return count
}
