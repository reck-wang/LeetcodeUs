package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(threeSum([]int{-1, 0, 1, 2, -1, -4}))
}

func threeSum(nums []int) [][]int {
	res := make(map[[3]int]struct{})
	sort.Ints(nums)

	for i, v := range nums[:len(nums)-2] {
		if i >= 1 && v == nums[i-1] {
			continue
		}

		dict := make(map[int]struct{})
		for _, x := range nums[i+1:] {
			if _, has := dict[x]; has {
				res[[3]int{v, -v - x, x}] = struct{}{}
			} else {
				dict[-v-x] = struct{}{}
			}
		}
	}

	var ans [][]int
	for tr := range res {
		ans = append(ans, tr[:])
	}

	return ans
}
