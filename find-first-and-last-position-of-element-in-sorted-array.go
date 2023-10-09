// https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	if len(nums) == 1 {
		if nums[0] != target {
			return []int{-1, -1}
		}
		return []int{0, 0}
	}

	left := 0
	right := len(nums) - 1
	var mid int

	for left <= right {
		mid = (left + right) / 2
		if target < nums[mid] {
			right = mid - 1
		} else if target > nums[mid] {
			left = mid + 1
		} else {
			break
		}
	}

	if nums[mid] != target {
		return []int{-1, -1}
	}
	
	left2 := left
	right2 := mid

	for left2 <= right2 {
		mid2 := (left2 + right2) / 2
		if target > nums[mid2] {
			left2 = mid2 + 1
		} else if target == nums[mid2] {
			right2 = mid2 - 1
		}
	}

	left3 := mid
	right3 := right

	for left3 <= right3 {
		mid3 := (left3 + right3) / 2
		if target == nums[mid3] {
			left3 = mid3 + 1
		} else if target < nums[mid3] {
			right3 = mid3 - 1
		}
	}

	return []int{left2, right3}
}
