// https://leetcode.com/problems/koko-eating-bananas
// ineffective, but works. 

func minEatingSpeed(piles []int, h int) int {
	var totalB int
	for _, bcount := range piles {
		totalB += bcount
	}

	minSpeed := totalB / h
	if minSpeed == 0 {
		minSpeed = 1
	}
	
	var totalHours int
	var hourCountEnough bool

	for !hourCountEnough {
		for _, bCount := range piles {
			totalHours = totalHours + bCount / minSpeed
			if bCount % minSpeed > 0 {
				totalHours++
			}
		}

		if totalHours <= h {
			return minSpeed
		}

		totalHours = 0
		minSpeed++
	}

	return 0
}
