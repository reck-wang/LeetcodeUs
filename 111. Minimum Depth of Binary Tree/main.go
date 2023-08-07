package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := minDepth(root.Left)
	rightDepth := minDepth(root.Right)
	if leftDepth == 0 || rightDepth == 0 {
		return 1 + leftDepth + rightDepth
	}
	return 1 + min(leftDepth, rightDepth)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
