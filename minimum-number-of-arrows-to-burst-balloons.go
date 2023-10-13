// https://leetcode.com/problems/minimum-number-of-arrows-to-burst-balloons


func findMinArrowShots(points [][]int) int {

	sort.Slice(points, func(i, j int) bool {
		if points[i][0] < points[j][0] {
			return true
		}
		if points[i][0] == points[j][0] {
			return points[i][1] < points[j][1]
		}
		return false
	})

	var pointer, start, end int
	var currentBalloon []int

	var arrowCount int

	for {
		if pointer == len(points) {
			break
		}

		currentBalloon = points[pointer]
		pointer++
		
		if start == 0 && end == 0 {
			start, end = currentBalloon[0], currentBalloon[1]
			continue
		}

		if currentBalloon[0] > end {
			arrowCount++
			start, end = currentBalloon[0], currentBalloon[1]
			continue
		}

		if currentBalloon[1] < end {
			end = currentBalloon[1]
			continue
		}
	}

	return arrowCount + 1
}
