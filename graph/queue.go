package graph

type Queue struct {
	nodes []Node
}

func NewQueue() *Queue {
	q := Queue{}
	q.nodes = []Node{}
	return &q
}

func (q *Queue) Enqueue(n Node) {
	q.nodes = append(q.nodes, n)
}

func (q *Queue) Dequeue() *Node {
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return &node
}

func (q *Queue) IsEmpty() bool {
	return len(q.nodes) == 0
}
