package dp

//LeetCode 121
func maxProfit(prices []int) int {
	bestPrice := prices[0]
	bestSell := 0
	curBestSell := 0

	for _, v := range prices {
		if v < bestPrice {
			bestPrice = v
		}
		curBestSell = v - bestPrice

		if curBestSell > bestSell {
			bestSell = curBestSell
		}

	}

	return bestSell
}
