// https://leetcode.com/problems/is-subsequence/

func isSubsequence(s string, t string) bool {
	var pointer int
	charList := []rune(s)
	if len(charList) == 0 {
		return true
	}
	
	for _, char := range t {
		if char == charList[pointer] {
			pointer++
			if pointer == len(charList) {
				return true
			}
		}
	}
	return false
}
