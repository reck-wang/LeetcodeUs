package main

func maxProfit(k int, prices []int) int {
	n := len(prices)
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, k+1)
	}

	for i := 0; i < k+1; i++ {
		dp[0][i][1] = -prices[0]
	}

	for i := 1; i < n; i++ {
		for j := 1; j < k+1; j++ {
			dp[i][j][0] = max(dp[i-1][j][0], dp[i-1][j][1]+prices[i])
			dp[i][j][1] = max(dp[i-1][j][1], dp[i-1][j-1][0]-prices[i])
		}
	}

	return dp[n-1][k][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
