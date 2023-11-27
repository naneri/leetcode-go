// https://leetcode.com/problems/knight-dialer

func knightDialer(n int) int {
	theNum := 1000000007
	theMap := make(map[int]int)

	theMap = map[int]int{
		0: 1,
		1: 1,
		2: 1,
		3: 1,
		4: 1,
		5: 1,
		6: 1,
		7: 1,
		8: 1,
		9: 1,
	}

	for i := 0; i < n-1; i++ {
		newMap := make(map[int]int)
		newMap[0] = theMap[6] + theMap[4]
		newMap[1] = theMap[8] + theMap[6]
		newMap[2] = theMap[9] + theMap[7]
		newMap[3] = theMap[4] + theMap[8]
		newMap[4] = theMap[0] + theMap[3] + theMap[9]
		newMap[6] = theMap[0] + theMap[1] + theMap[7]
		newMap[7] = theMap[2] + theMap[6]
		newMap[8] = theMap[1] + theMap[3]
		newMap[9] = theMap[2] + theMap[4]

		for key, val := range newMap {
			newMap[key] = val % theNum
		}
		
		theMap = newMap
	}
	
	var answer int
	for _, val := range theMap {
		answer += val
	}

	return answer % theNum
}
