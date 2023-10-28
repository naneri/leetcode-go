// https://leetcode.com/problems/count-vowels-permutation/

func countVowelPermutation(n int) int {
	answer := dp(n)
	var total int
	for _, val := range answer {
		total += val
	}

	return total % modulo
}

var modulo int = 1000000007
var defaultMap = map[rune]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}

func dp(n int) map[rune]int {
	if n == 1 {
		return defaultMap
	}

	prev := dp(n - 1)

	res := map[rune]int{
		'a': 0, 'e': 0, 'i': 0, 'o': 0, 'u': 0,
	}

	res['a'] = (prev['u'] + prev['e'] + prev['i']) % modulo
	res['e'] = (prev['a'] + prev['i']) % modulo
	res['i'] = (prev['e'] + prev['o']) % modulo
	res['o'] = prev['i'] % modulo
	res['u'] = (prev['i'] + prev['o']) % modulo

	return res
}
