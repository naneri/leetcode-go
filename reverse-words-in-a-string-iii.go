// https://leetcode.com/problems/reverse-words-in-a-string-iii

func reverseWords(s string) string {
	stringList := strings.Split(s, " ")

	for key, word := range stringList {
		stringList[key] = reverseString(word)
	}
	
	return strings.Join(stringList, " ")
}

func reverseString(s string) string {
	runeSlice := []rune(s)
	
	for i := 0; i < len(runeSlice) / 2; i++ {
		runeSlice[i], runeSlice[len(runeSlice) - i - 1] =  runeSlice[len(runeSlice) - i - 1], runeSlice[i]
	}
	
	return string(runeSlice)
}
