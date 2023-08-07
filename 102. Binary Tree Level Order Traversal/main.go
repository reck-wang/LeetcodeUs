package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// method one: use BFS
//func levelOrder(root *TreeNode) [][]int {
//	var ans [][]int
//
//	if root == nil {
//		return ans
//	}
//
//	var queue = []*TreeNode{root}
//
//	for len(queue) > 0 {
//		tempLen := len(queue)
//		level := make([]int, 0, tempLen)
//		for i := 0; i < tempLen; i++ {
//			pop := queue[0]
//			queue = queue[1:]
//			if pop.Left != nil {
//				queue = append(queue, pop.Left)
//			}
//			if pop.Right != nil {
//				queue = append(queue, pop.Right)
//			}
//
//			level = append(level, pop.Val)
//		}
//		ans = append(ans, level)
//	}
//
//	return ans
//}
