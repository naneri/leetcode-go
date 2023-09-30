// https://leetcode.com/problems/132-pattern/submissions/

type IncreaseInterval struct {
	End      int
	Interval []int
}

func find132pattern(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	var end int
	allIntervals := make([]IncreaseInterval, 0)
	interval := make([]int, 0)

	for i := 0; i < len(nums); i++ {

		if len(interval) == 0 {
			interval = append(interval, nums[i])
			continue
		}

		if len(interval) == 1 {
			if nums[i] < interval[0] {
				interval[0] = nums[i]
			}

			if nums[i] > interval[0] {
				end = i
				interval = append(interval, nums[i])
			}
		}

		if len(interval) == 2 {
			if nums[i] < interval[0] {
			if len(allIntervals) > 0 {
					if interval[0] <= allIntervals[len(allIntervals)-1].Interval[0] {
						if interval[1] >= allIntervals[len(allIntervals)-1].Interval[1] {
							allIntervals[len(allIntervals)-1].Interval[0] = interval[0]
							allIntervals[len(allIntervals)-1].Interval[1] = interval[1]
							allIntervals[len(allIntervals)-1].End = end
							interval = make([]int, 0)
							interval = append(interval, nums[i])
							end = 0
							continue
						}
					}
				}
				allIntervals = append(allIntervals, IncreaseInterval{
					End:      end,
					Interval: interval,
				})
				interval = make([]int, 0)
				interval = append(interval, nums[i])
				end = 0
				continue
			}

			if nums[i] > interval[1] {
				end = i
				interval[1] = nums[i]
			}

			if nums[i] > interval[0] && nums[i] < interval[1] {
				return true
			}
		}
	}

	if len(interval) == 2 {
		allIntervals = append(allIntervals, IncreaseInterval{
			End:      end,
			Interval: interval,
		})
	}
	fmt.Println(allIntervals)
	for i := len(nums) - 1; i >= 0; i-- {
		number := nums[i]

		for j := len(allIntervals) - 1; j >= 0; j-- {
			if i <= allIntervals[j].End {
				continue
			}

			if number > allIntervals[j].Interval[0] && number < allIntervals[j].Interval[1] {
				return true
			}
		}
	}

	return false
}
