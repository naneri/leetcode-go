// https://leetcode.com/problems/sort-array-by-parity

func sortArrayByParity(nums []int) []int {
	lptr := 0
	rptr := len(nums) - 1
	for lptr < rptr {
		
		if nums[lptr] & 1  == 0{
			lptr++
			continue
		}

		if nums[rptr] & 1 ==  1 {
			rptr--
			continue
		}
		nums[lptr], nums[rptr] = nums[rptr], nums[lptr]
		lptr++
		rptr--
	}

	return nums
}
