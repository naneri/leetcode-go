// https://leetcode.com/problems/unique-paths

var cache map[string]int 

func uniquePaths(m int, n int) int {
	cache = make(map[string]int)

	return getPathLength(m, n)
}

func getPathLength(m int, n int) int {
	code := fmt.Sprintf("%d-%d", m, n)

	if val, ok := cache[code]; ok {
		return val
	}
	if m == 2 && n == 1 {
		return 1
	}
	if m == 1 && n == 2 {
		return 1
	}
	if m == 1 && n == 1 {
		return 1
	}

	var total int

	if m > 1 {
		total += getPathLength(m-1, n)
	}

	if n > 1 {
		total += getPathLength(m, n-1)
	}

	cache[code] = total
	return total
}
