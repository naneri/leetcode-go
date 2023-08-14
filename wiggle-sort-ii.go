// https://leetcode.com/problems/wiggle-sort-ii

func wiggleSort(nums []int)  {
	if len(nums) < 2 {
		return
	}

	sort.Ints(nums)

	var result []int
	var reducer int

	if len(nums) % 2 == 0 {
		reducer = 1
	}

	for i := 0; i < len(nums) / 2; i++ {
		result = append(result, nums[len(nums) / 2 - i - reducer])
		result = append(result, nums[len(nums) - i - 1])
	}


	if len(result) < len(nums) {
		result = append(result, nums[0])
	}

	for key := range nums {
		nums[key] = result[key]
	}
}
