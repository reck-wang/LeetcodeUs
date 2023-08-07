package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// iterative
//func reverseKGroup(head *ListNode, k int) *ListNode {
//	length := 0
//
//	for temp := head; temp != nil; temp = temp.Next {
//		length++
//	}
//
//	dummyHead := &ListNode{Next: head}
//	prev := dummyHead
//
//	for length >= k {
//		cur := prev.Next
//		for i := 1; i < k; i++ {
//			next := cur.Next
//			cur.Next = next.Next
//			next.Next = prev.Next
//			prev = next
//			fmt.Println("curent node: ", cur)
//			fmt.Println("next node: ", next)
//			fmt.Println("prev node: ", prev)
//		}
//		fmt.Println()
//		prev = cur
//		length -= k
//	}
//
//	return dummyHead.Next
//}

// recursive
func reverseKGroup(head *ListNode, k int) *ListNode {
	l := 0
	for temp := head; temp != nil; temp = temp.Next {
		l++
	}

	return reverse(head, l, k)
}

func reverse(node *ListNode, l, k int) *ListNode {
	if l < k {
		return node
	}

	var prev, next *ListNode
	cur := node

	for i := 0; i < k; i++ {
		next = cur.Next
		cur.Next = prev
		prev = cur
		cur = next
	}

	node.Next = reverse(next, l-k, k)
	return prev
}
