package main

func twoSum(nums []int, target int) []int {
	set := map[int]int{}

	for i, num := range nums {
		temp := target - num
		if _, has := set[temp]; has {
			return []int{set[temp], i}
		}

		set[num] = i
	}

	return []int{-1, -1}
}
