package main

type Stack struct {
	data []int
}

func newStack(s *Stack) *Stack {
	return &Stack{data: make([]int, 0)}
}

func (s *Stack) Push(a int) {
	s.data = append(s.data, a)
}

func (s *Stack) Pop() (int, bool) {
	if len(s.data) == 0 {
		return 0, false
	}
	topIndex := len(s.data) - 1
	val := s.data[topIndex]
	s.data = s.data[:topIndex]
	return val, true
}

func (s *Stack) isEmpty() bool {
	return len(s.data) == 0
}