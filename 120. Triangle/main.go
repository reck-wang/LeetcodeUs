package main

func minimumTotal(triangle [][]int) int {
	m := len(triangle)
	if m == 1 {
		return triangle[0][0]
	}

	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, i+1)
	}

	dp[0][0] = triangle[0][0]
	for i := 1; i < m; i++ {
		for j := 0; j < len(dp[i]); j++ {
			if j != 0 && j != len(dp[i])-1 {
				dp[i][j] = min(dp[i-1][j-1], dp[i-1][j]) + triangle[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][0] + triangle[i][j]
			} else {
				dp[i][j] = dp[i-1][j-1] + triangle[i][j]
			}
		}
	}

	ans := dp[m-1][0]
	for i := 1; i < len(dp[m-1]); i++ {
		if dp[m-1][i] < ans {
			ans = dp[m-1][i]
		}
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
