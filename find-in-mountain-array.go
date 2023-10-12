// https://leetcode.com/problems/find-in-mountain-array

func findInMountainArray(target int, mountainArr *MountainArray) int {
	if mountainArr.length() < 100 {
		for i := 0; i < mountainArr.length(); i++ {
			if mountainArr.get(i) == target {
				return i
			}
		}
		return -1
	}

	arrayLen := mountainArr.length()

	left := 0
	right := arrayLen - 1

	// 1,3,5,4,2
	// 2 (0,4)| 3,5
	// 3 (2,4)| 5,4
	// 2 (2,3)| 3,5
	// 3 (3,3)| 5,4


	for left < right {
		mid := (left + right) / 2
		midVal := mountainArr.get(mid)
		if mountainArr.get(mid - 1) < midVal  {
			left = mid
			if right - left < 3 {
				if mountainArr.get(mid + 1) < midVal {
					left = mid
					break
				}
			}
		} else {
			right = mid
		}

	}

	top := left

	left = 0
	right = top

	for left <= right {
		mid := (left + right) / 2

		arrVal := mountainArr.get(mid)
		if target < arrVal {
			right = mid-1
		} else if target > arrVal {
			left = mid+1
		} else {
			return mid
		}
	}

	left = top
	right = arrayLen - 1

	for left <= right {
		mid := (left + right) / 2
		arrVal := mountainArr.get(mid)
		if target < arrVal   {
			left = mid+1
		} else if target > arrVal   {
			right = mid-1
		} else {
			return mid
		}
	}

	return -1
}
