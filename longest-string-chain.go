// https://leetcode.com/problems/longest-string-chain

func longestStrChain(words []string) int {
	var answer int

	results := make(map[string]int)

	sort.Slice(words, func(i, j int) bool {
		return len(words[i]) < len(words[j])
	})

	for _, word := range words {
		var maxLength int
		for index := range word {
			newWord := word[:index] + word[index+1:]
			if _, ok := results[newWord]; ok {
				if maxLength < results[newWord] + 1 {
					maxLength = results[newWord]+1
				}
			}
		}

		results[word] = maxLength
		if maxLength > answer {
			answer = maxLength
		}
	}
	
	return answer + 1
}
