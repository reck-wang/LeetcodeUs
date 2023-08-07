package main

import "math"

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

// use dfs
//func isValidBST(root *TreeNode) bool {
//	return isValid(root, math.MinInt32, math.MaxInt32)
//}
//
//func isValid(node *TreeNode, min, max int) bool {
//	if node == nil {
//		return true
//	}
//
//	if node.Val <= min || node.Val >= max {
//		return false
//	}
//
//	return isValid(node.Left, min, node.Val) &&
//		isValid(node.Right, node.Val, max)
//}

// use binary search tree's inOrder
func isValidBST(root *TreeNode) bool {
	var inOrder func(*TreeNode) bool
	prev := math.MinInt32 - 1

	inOrder = func(node *TreeNode) bool {
		if node == nil {
			return true
		}

		if !inOrder(node.Left) {
			return false
		}
		if prev >= node.Val {
			return false
		}
		prev = node.Val

		return inOrder(node.Right)
	}

	return inOrder(root)
}
