// https://leetcode.com/problems/candy/

type Higher struct {
	left  int
	right int
}

func candy(ratings []int) int {

	progression := make([]Higher, 0)

	for key := range ratings {
		progression = append(progression, Higher{})
		if key != 0 {
			if ratings[key] > ratings[key-1] {
				progression[key].left = progression[key-1].left + 1
			}
		}
	}

	for i := len(ratings) - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] {
			progression[i].right = progression[i+1].right + 1
		}
	}
	
	var result int
	for _, val := range progression {
		result += maxInt(val.left, val.right)
	}
	return result + len(ratings)
}

func maxInt(a, b int) int {
	if a > b {
		return a
	}
	return b
}
