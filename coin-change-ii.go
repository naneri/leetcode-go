// https://leetcode.com/problems/coin-change-ii/



func change(amount int, coins []int) int {

	// reflects variants for all values up to amount. Like 0,1,2,3,4,5,6,7
	combs := make([]int, amount+1)

	// 0 amount can be created by a combination of 0 coins. So there is only a single variant on how to gather such amount
	combs[0] = 1


	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			combs[i] += combs[i - coin]
		}
	}

	return combs[amount]
}
