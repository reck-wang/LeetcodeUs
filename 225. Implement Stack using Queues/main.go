package main

type MyStack struct {
	queue1 []int
	queue2 []int
}

func Constructor() MyStack {
	return MyStack{
		make([]int, 0),
		make([]int, 0),
	}
}

func (s *MyStack) Push(x int) {
	if len(s.queue1) != 0 {
		s.queue1 = append(s.queue1, x)
	}
	s.queue2 = append(s.queue2, x)
}

func (s *MyStack) Pop() int {
	if len(s.queue1) != 0 {
		l := len(s.queue1) - 1
		res := s.queue1[l]
		for i := range s.queue1[:l] {
			s.queue2 = append(s.queue2, s.queue1[i])
		}
		s.queue1 = s.queue1[:0]

		return res
	}

	l := len(s.queue2) - 1
	res := s.queue2[l]
	for i := range s.queue2[:l] {
		s.queue1 = append(s.queue1, s.queue2[i])
	}
	s.queue2 = s.queue2[:0]

	return res

}

func (s *MyStack) Top() int {
	if len(s.queue1) != 0 {
		return s.queue1[len(s.queue1)-1]
	}
	return s.queue2[len(s.queue2)-1]
}

func (s *MyStack) Empty() bool {
	return len(s.queue1) == 0 && len(s.queue2) == 0
}
