package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// recursion
//func reverseList(node *ListNode) *ListNode {
//	if node == nil || node.Next == nil {
//		return node
//	}
//
//	newNode := reverseList(node.Next)
//	node.Next.Next = node
//	node.Next = nil
//
//	return newNode
//}

// replace head
func reverseList(head *ListNode) *ListNode {
	prev := &ListNode{}
	temp := head

	for temp != nil {
		next := temp.Next
		temp.Next = prev.Next
		prev.Next = temp
		temp = next
	}

	return prev.Next
}
