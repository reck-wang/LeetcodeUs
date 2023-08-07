package main

import (
	"container/heap"
	"sort"
)

type KthLargest struct {
	sort.IntSlice
	size int
}

func Constructor(k int, nums []int) KthLargest {
	hp := &KthLargest{
		IntSlice: nums,
		size:     k,
	}

	heap.Init(hp)

	return *hp
}

func (k *KthLargest) Add(val int) int {
	heap.Push(k, val)
	for k.Len() > k.size {
		heap.Pop(k)
	}

	return k.IntSlice[0]
}

func (k *KthLargest) Push(x any) {
	k.IntSlice = append(k.IntSlice, x.(int))
}

func (k *KthLargest) Pop() any {
	res := k.IntSlice[len(k.IntSlice)-1]
	k.IntSlice = k.IntSlice[:len(k.IntSlice)-1]
	return res
}
