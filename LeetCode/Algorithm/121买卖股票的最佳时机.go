package algorithm

func maxProfit(prices []int) int {
	minPrice := prices[0]
	maxMonry := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		if price-minPrice>maxMonry{
			maxMonry = price-minPrice
		}
	}
	return maxMonry
}
