//  https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

func maxProfit(prices []int) int {
	var profit int
	minPrice := prices[0]

	for _, val := range prices {
		if val < minPrice {
			minPrice = val
		} else if (val - minPrice) > profit {
			profit = val - minPrice
		}
	}

	return profit
}
