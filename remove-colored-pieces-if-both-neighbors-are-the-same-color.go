// https://leetcode.com/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color


func winnerOfGame(colors string) bool {
	colorSlice := []rune(colors)

	var aTurns, bTurns int
	for i := 1; i < len(colorSlice) - 1; i++ {
		if colorSlice[i] == 'A' {
			if colorSlice[i - 1] == 'A' && colorSlice[i + 1] == 'A' {
				aTurns++
			}
		} else {
			if colorSlice[i - 1] == 'B' && colorSlice[i + 1] == 'B' {
				bTurns++
			}
		}
	}
	
	return aTurns > bTurns
}
