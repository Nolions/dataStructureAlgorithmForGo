package graph

type Stack struct {
	nodes []Node
}

func NewStack() *Stack {
	s := Stack{}
	s.nodes = []Node{}
	return &s
}

func (s *Stack) Push(n Node) {
	s.nodes = append(s.nodes, n)
}

func (s *Stack) Pop() *Node {
	node := s.nodes[len(s.nodes)-1]
	s.nodes = s.nodes[:len(s.nodes)-1]

	return &node
}

func (s *Stack) IsEmpty() bool {
	return len(s.nodes) == 0
}
