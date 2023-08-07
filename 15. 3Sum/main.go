package main

import "sort"

// method one: use set
//func threeSum(nums []int) [][]int {
//	res := make(map[[3]int]struct{})
//	sort.Ints(nums)
//
//	for i, v := range nums[:len(nums)-2] {
//		if i >= 1 && v == nums[i-1] {
//			continue
//		}
//
//		dict := make(map[int]struct{})
//		for _, x := range nums[i+1:] {
//			if _, has := dict[x]; has {
//				res[[3]int{v, -v - x, x}] = struct{}{}
//			} else {
//				dict[-v-x] = struct{}{}
//			}
//		}
//	}
//
//	var ans [][]int
//	for tr := range res {
//		temp := tr
//		ans = append(ans, temp[:])
//	}
//
//	return ans
//}

// method two: use two pointers
func threeSum(nums []int) [][]int {
	ans := [][]int{}
	sort.Ints(nums)

	for i, num := range nums[:len(nums)-2] {
		if i >= 1 && num == nums[i-1] {
			continue
		}

		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] < -num {
				l++
			} else if nums[l]+nums[r] > -num {
				r--
			} else {
				ans = append(ans, []int{num, nums[l], nums[r]})
				for l < r && nums[l] == nums[l+1] {
					l++
				}
				for l < r && nums[r] == nums[r-1] {
					r--
				}
				l++
				r--
			}
		}
	}

	return ans
}
