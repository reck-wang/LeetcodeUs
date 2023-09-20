package main

func maxProduct(nums []int) int {
	maxP := nums[0]
	minP := nums[0]
	ans := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxP, minP = minP, maxP
		}
		maxP = max(nums[i], nums[i]*maxP)
		minP = min(nums[i], nums[i]*minP)
		ans = max(ans, maxP)
	}

	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
