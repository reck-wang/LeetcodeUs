package main

type ListNode struct {
	Val  int
	Next *ListNode
}

// use hashset
//func hasCycle(head *ListNode) bool {
//	set := map[*ListNode]struct{}{}
//
//	for temp := head; temp != nil; temp = temp.Next {
//		if _, has := set[temp]; has {
//			return true
//		}
//		set[temp] = struct{}{}
//	}
//
//	return false
//}

// use fast and slow pointers
func hasCycle(head *ListNode) bool {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			return true
		}
	}

	return false
}
