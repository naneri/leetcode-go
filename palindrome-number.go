// https://leetcode.com/problems/palindrome-number/


func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	var intSlice []int

	for x > 0 {
		intSlice = append(intSlice, x % 10)
		x = x / 10
	}

	for i := 0; i < len(intSlice) / 2; i ++ {
		if intSlice[i] != intSlice[len(intSlice) - i - 1] {
			return false
		}
	}

	return true
}
