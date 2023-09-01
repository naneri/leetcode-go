// https://leetcode.com/problems/best-time-to-buy-and-sell-stock-ii/

func maxProfit(prices []int) int {
	currentMin := prices[0]
	var maximumProfit int
	var totalProfit int

	for _, val := range prices {

		if val < currentMin {
			if maximumProfit != 0 {
				totalProfit += maximumProfit
				maximumProfit = 0
			}
			currentMin = val
		} else {
			curProfit := val - currentMin

			if curProfit >= maximumProfit {
				maximumProfit = curProfit
				fmt.Println(val, "here")
			} else {
				fmt.Println(val, "or here")
				totalProfit 	+= maximumProfit
				maximumProfit 	= 0
				currentMin 		= val
			}
		}
		//fmt.Println(currentMin, maximumProfit, totalProfit)
	}

	if maximumProfit > 0 {
		return totalProfit + maximumProfit
	}

	return totalProfit
}
