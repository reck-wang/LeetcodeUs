package main

func countBits(n int) []int {
	counts := make([]int, n+1)
	for i := 1; i <= n; i++ {
		counts[i] = counts[i&(i-1)] + 1
	}
	return counts
}
