// https://leetcode.com/problems/find-the-difference

func findTheDifference(s string, t string) byte {

	sByte := []rune(s)
	tByte := []rune(t)

	sort.Slice(sByte, func(i, j int) bool {
		return sByte[i] < sByte[j]
	})

	sort.Slice(tByte, func(i, j int) bool {
		return tByte[i] < tByte[j]
	})

	for key := range sByte {
		if sByte[key] != tByte[key] {
			return byte(tByte[key])
		}
	}

	return byte(tByte[len(tByte)-1])
}
