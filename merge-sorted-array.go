//  https://leetcode.com/problems/merge-sorted-array

func merge(nums1 []int, m int, nums2 []int, n int) {
	var firstP, secondP int
	var newArr []int

	for {
		if len(newArr) == m+n {
			break
		}

		if firstP == m {
			newArr = append(newArr, nums2[secondP])
			secondP++
			continue
		}

		if secondP == n {
			newArr = append(newArr, nums1[firstP])
			firstP++
			continue
		}

		if nums1[firstP] <= nums2[secondP] {
			newArr = append(newArr, nums1[firstP])
			firstP++
		} else {
			newArr = append(newArr, nums2[secondP])
			secondP++
		}
	}

	for key, val := range newArr {
		nums1[key] = val
	}
}
