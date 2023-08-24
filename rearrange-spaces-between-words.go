// https://leetcode.com/problems/rearrange-spaces-between-words

func reorderSpaces(text string) string {

	var spaceCount int
	var wordCount  int
	var newWord []rune
	words := strings.Fields(text)
	
	for _, char := range text {
		if char == ' ' {
			spaceCount++
		}
	}

	if len(newWord) > 0 {
		anotherWord := string(newWord)
		words = append(words, anotherWord)
		newWord = make([]rune, 0)
		wordCount++
	}

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaceCount)
	}

	separator := strings.Repeat(" ", spaceCount/(len(words)-1))
	padding := strings.Repeat(" ", spaceCount%(len(words)-1))

	return strings.Join(words, separator) + padding

}
