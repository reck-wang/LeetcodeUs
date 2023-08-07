package main

type MyQueue struct {
	stack1 []int
	stack2 []int
}

func Constructor() MyQueue {
	return MyQueue{
		stack1: make([]int, 0),
		stack2: make([]int, 0),
	}
}

func (q *MyQueue) Push(x int) {
	q.stack1 = append(q.stack1, x)
}

func (q *MyQueue) Pop() int {
	if len(q.stack2) == 0 {
		for len(q.stack1) != 0 {
			q.stack2 = append(q.stack2, q.stack1[len(q.stack1)-1])
			q.stack1 = q.stack1[:len(q.stack1)-1]
		}
	}

	result := q.stack2[len(q.stack2)-1]
	q.stack2 = q.stack2[:len(q.stack2)-1]
	return result
}

func (q *MyQueue) Peek() int {
	if len(q.stack2) != 0 {
		return q.stack2[len(q.stack2)-1]
	} else {
		return q.stack1[0]
	}
}

func (q *MyQueue) Empty() bool {
	return len(q.stack1) == 0 && len(q.stack2) == 0
}
