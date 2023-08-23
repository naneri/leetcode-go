// https://leetcode.com/problems/reorganize-string

type charStr struct {
	char rune
	count int
}

func reorganizeString(s string) string {
	stringMap := make(map[rune]int)
	charSlice := make([]charStr, 0)

	for _, char := range s {
		stringMap[char]++
	}

	for key, val := range stringMap {
		charSlice = append(charSlice, charStr{
			char:  key,
			count: val,
		})
	}

	sort.Slice(charSlice, func(i, j int) bool {
		return charSlice[i].count > charSlice[j].count
	})

	if charSlice[0].count > (len(s) + 1) / 2  {
		return ""
	}

	res := make([]rune, len(s))
  
	var index int
  
	for _, ch := range charSlice {
		for i := 0; i < ch.count; i++ {
			if index >= len(res) {
				index = 1
			}
			res[index] = ch.char
			index = index + 2
		}
	}

	return string(res)
}
