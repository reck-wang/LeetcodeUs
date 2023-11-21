package main

func maxProfit(prices []int) int {
	n := len(prices)
	dp := make([][2][2]int, n)

	dp[0][0][1] = -prices[0]

	for i := 1; i < n; i++ {
		dp[i][0][0] = max(dp[i-1][0][0], dp[i-1][1][0])
		dp[i][1][0] = dp[i-1][0][1] + prices[i]
		dp[i][0][1] = max(dp[i-1][0][1], dp[i-1][0][0]-prices[i])
	}

	var ans int
	for i := 0; i < n; i++ {
		for j := 0; j < 2; j++ {
			if ans < dp[i][j][0] {
				ans = dp[i][j][0]
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
