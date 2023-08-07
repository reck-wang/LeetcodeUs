package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{Next: head}

	pre := dummyHead
	temp := head
	for temp != nil && temp.Next != nil {
		pre.Next = temp.Next
		temp.Next = pre.Next.Next
		pre.Next.Next = temp

		pre = temp
		temp = temp.Next
	}

	return dummyHead.Next
}
