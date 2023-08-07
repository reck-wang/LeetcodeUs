package main

// greedy strategy
func maxProfit(prices []int) int {
	profit := 0

	for i, price := range prices[:len(prices)-1] {
		if price < prices[i+1] {
			profit += prices[i+1] - price
		}
	}

	return profit
}
