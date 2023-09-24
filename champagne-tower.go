// https://leetcode.com/problems/champagne-tower/

func champagneTower(poured int, query_row int, query_glass int) float64 {
	if poured == 0 {
		return 0
	}
	var currentRow int
	glasses := make(map[int]map[int]float64)
	pouredFloat64 := float64(poured)

	var keepIterating bool
	glasses[0] = make(map[int]float64)
	glasses[0][0] = pouredFloat64
	if pouredFloat64 > 1 {
		keepIterating = true
	}
	currentRow++

	totalGlasses := 1
	for keepIterating {
		keepIterating = false
		glasses[currentRow] = make(map[int]float64)

		for i := 0; i < currentRow+1; i++ {
			totalGlasses++
			var totalExcess float64
			if i != currentRow {
				rightGlass := glasses[currentRow-1][i]
				if rightGlass > 1 {
					totalExcess += (rightGlass - 1) / 2
				}
			}

			if i != 0 {
				leftGlass := glasses[currentRow-1][i-1]
				if leftGlass > 1 {
					totalExcess += (leftGlass - 1) / 2
				}
			}
			if totalExcess > 1 {
				keepIterating = true
			}
			glasses[currentRow][i] = totalExcess

		}

		currentRow++
		if currentRow > query_row {
			break
		}
	}

	if query_row > currentRow {
		return 0
	}

	glass := glasses[query_row][query_glass]

	if glass > 1 {
		return 1
	}

	return glass
}
