// https://leetcode.com/problems/rearrange-spaces-between-words

func reorderSpaces(text string) string {

	var spaceCount int
	words := strings.Fields(text)
	
	for _, char := range text {
		if char == ' ' {
			spaceCount++
		}
	}

	if len(words) == 1 {
		return words[0] + strings.Repeat(" ", spaceCount)
	}

	separator := strings.Repeat(" ", spaceCount/(len(words)-1))
	padding := strings.Repeat(" ", spaceCount%(len(words)-1))

	return strings.Join(words, separator) + padding

}
