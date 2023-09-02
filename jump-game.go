// https://leetcode.com/problems/jump-game

func canJump(nums []int) bool {
	var currentMaxDist int

	for i := 0; i < len(nums); i++ {
		if i == len(nums) - 1  {
			return true
		}
		currentJumpDist := nums[i] + i
		if currentJumpDist > currentMaxDist {
			currentMaxDist = currentJumpDist
		}
		
		if nums[i] == 0 && currentMaxDist == i {
			return false
		}
	}
	return true
}
