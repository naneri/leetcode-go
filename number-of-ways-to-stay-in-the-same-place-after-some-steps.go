// https://leetcode.com/problems/number-of-ways-to-stay-in-the-same-place-after-some-steps

var cache = make(map[string]int)

const MOD = 1e9 + 7

func numWays(steps int, arrLen int) int {
	return stepsFromPos(steps, arrLen, 0)
}

func stepsFromPos(steps int, arrLen int, pos int) int {
	var result int
	key := fmt.Sprintf("%d-%d-%d", steps, arrLen, pos)

	if val, ok := cache[key]; ok {
		return val
	}

	if pos < 0 {
		return 0
	}

	if pos > steps {
		return 0
	}

	if pos >= arrLen {
		return 0
	}

	if steps == 0 {
		if pos == 0 {
			return 1
		} else {
			return 0
		}
	}

	result += stepsFromPos(steps-1, arrLen, pos-1)
	result += stepsFromPos(steps-1, arrLen, pos)
	result += stepsFromPos(steps-1, arrLen, pos+1)

	result = result % MOD
	
	cache[key] = result
	
	return result
}
