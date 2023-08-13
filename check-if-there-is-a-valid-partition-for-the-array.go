// https://leetcode.com/problems/check-if-there-is-a-valid-partition-for-the-array/


func validPartition(nums []int) bool {
	dp := make([]bool, len(nums)+1)
	dp[0] = true


	for i := 0; i < len(nums); i++ {
		dpIndex := i + 1

		if i > 0 && nums[i] == nums[i-1] {
			dp[dpIndex] = dp[dpIndex] || dp[dpIndex-2]
		}
		if i > 1 && (nums[i] == nums[i-1]) && (nums[i] == nums[i-2]) {
			dp[dpIndex] = dp[dpIndex] || dp[dpIndex-3]
		}
		if i > 1 && (nums[i] == nums[i-1] + 1) && (nums[i] == nums[i-2] + 2) {
			dp[dpIndex] = dp[dpIndex] || dp[dpIndex-3]
		}
	}

	return dp[len(nums)]
}
