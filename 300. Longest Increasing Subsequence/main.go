package main

// Method One: Dynamic Programming
//func lengthOfLIS(nums []int) int {
//	n := len(nums)
//	if n == 1 {
//		return 1
//	}
//
//	dp := make([]int, n)
//	for i := range dp {
//		dp[i] = 1
//	}
//
//	for i := 0; i < n; i++ {
//		for j := 0; j < i; j++ {
//			if nums[i] > nums[j] {
//				dp[i] = max(dp[i], dp[j]+1)
//			}
//		}
//	}
//
//	var ans int
//	for _, v := range dp {
//		if v > ans {
//			ans = v
//		}
//	}
//	return ans
//}
//
//func max(a, b int) int {
//	if a > b {
//		return a
//	}
//	return b
//}

// Method Two: Binary Search
func lengthOfLIS(nums []int) int {
	n := len(nums)
	if n == 1 {
		return 1
	}

	lis := []int{nums[0]}
	for i := 1; i < n; i++ {
		idx := binarySearch(lis, nums[i])

		if idx >= 0 && idx < len(lis) && nums[i] < lis[idx] {
			lis[idx] = nums[i]
		} else if idx == len(lis) {
			lis = append(lis, nums[i])
		}
	}

	return len(lis)
}

func binarySearch(arr []int, target int) int {
	l, r := 0, len(arr)
	for l < r {
		mid := l + (r-l)/2
		if arr[mid] < target {
			l = mid + 1
		} else {
			r = mid
		}
	}
	return l
}
