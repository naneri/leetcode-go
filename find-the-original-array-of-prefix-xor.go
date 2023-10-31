// https://leetcode.com/problems/find-the-original-array-of-prefix-xor

func findArray(pref []int) []int {
	var res []int

	var curNum int

	for _, val := range pref {
		res = append(res, val ^ curNum)
		curNum = val
	}

	return res
}
