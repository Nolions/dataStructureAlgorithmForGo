package array

import "github.com/Nolions/dataStructureAlgorithmForGo/stack"

type Stack struct {
	list []stack.Data
}

func New(size int) *Stack {
	d := new(Stack)
	d.list = make([]stack.Data, 0, size)

	return d
}

func (s *Stack) Push(data stack.Data) {
	s.list = append(s.list, data)
}

func (s *Stack) Pop() stack.Data {
	tmp := s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]

	return tmp
}

func (s *Stack) isEmpty() bool {
	return len(s.list) > 0
}
