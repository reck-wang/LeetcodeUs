package main

func coinChange(coins []int, amount int) int {
	max := amount + 1
	dp := make([]int, max)
	dp[0] = 0
	for i := 1; i < max; i++ {
		dp[i] = max
	}

	for i := 1; i <= amount; i++ {
		for _, coin := range coins {
			if coin <= i {
				dp[i] = min(dp[i], dp[i-coin]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
