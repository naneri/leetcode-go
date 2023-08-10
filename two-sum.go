// https://leetcode.com/problems/two-sum/

func twoSum(nums []int, target int) []int {
	myMap :=  make(map[int]int)

	for key, val := range nums {
		myMap[val] = key
	}

	for key, val := range nums {
		searchKey := target - val
		if res, ok := myMap[searchKey]; ok {
			if key != res {
				return []int{key, res}
			}

		}
	}

	return nil
}
