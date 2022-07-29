package dp

//My Original Work
func maxProfit2(prices []int) int {
	bestPrice := prices[0]
	bestSell := 0
	curBestSell := 0
	profit := 0
	curBestProfit := 0

	for _, v := range prices {

		if v < bestPrice {
			bestPrice = v
		}

		curBestSell = v - bestPrice

		if curBestSell > bestSell {
			bestSell = curBestSell
			curBestProfit = bestSell
			curBestSell = 0
		}
		if curBestProfit+curBestSell > profit {
			profit = curBestProfit + curBestSell
			bestPrice = profit
		}

	}
	return profit
}

//YT Inspiration
func maxProfit2YT(prices []int) int {
	profit := 0

	for i := range prices {
		if i == 0 {
			continue
		}

		//compare current with previous
		if prices[i] > prices[i-1] {
			profit += (prices[i] - prices[i-1])
		}
	}

	return profit
}
