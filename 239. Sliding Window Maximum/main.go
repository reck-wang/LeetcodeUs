package main

// use dequeue
func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	var queue, ans []int

	for i, num := range nums {
		// 移除超出左边界的数据
		if i >= k && queue[0] <= i-k {
			queue = queue[1:]
		}

		// 移除比新加入的数据小的数据
		for len(queue) > 0 && nums[queue[len(queue)-1]] <= num {
			queue = queue[:len(queue)-1]
		}
		// 加入新的索引
		queue = append(queue, i)

		if i >= k-1 {
			// 最大的值在最左边
			ans = append(ans, nums[queue[0]])
		}
	}

	return ans
}
