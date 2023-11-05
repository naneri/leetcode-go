// https://leetcode.com/problems/find-the-winner-of-an-array-game

func getWinner(arr []int, k int) int {
	if k > len(arr) {
		var maxVal int
		for _, val := range arr {
			if val > maxVal {
				maxVal = val
			}
		}
		return maxVal
	}
	
	var round int
	var maxVal int
	
	if arr[0] > arr[1] {
		maxVal = arr[0]
	} else {
		maxVal = arr[1]
	}
	
	for key, val := range arr {
		if key == 0 {
			continue
		}	
		round++
		if val > maxVal {
			maxVal = val
			round = 1
		}	
		if round == k {
			return maxVal
		}
	}
	return maxVal
}
