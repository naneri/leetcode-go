// https://leetcode.com/problems/excel-sheet-column-title

func convertToTitle(columnNumber int) string {
	var nums []int
	for columnNumber > 0 {
		fnum := (columnNumber - 1) % 26
		columnNumber = (columnNumber - 1) / 26
		nums = append(nums, fnum)
	}

	var res string
	for i := len(nums) - 1; i >=0; i--{
		res = res + string(rune(nums[i] + 65))
	}
	
	return res
}
