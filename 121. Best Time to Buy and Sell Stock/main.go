package main

func maxProfit(prices []int) int {
	var ans, min int = 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < min {
			min = prices[i]
		} else if prices[i] > min {
			if prices[i]-min > ans {
				ans = prices[i] - min
			}
		}
	}
	return ans
}
