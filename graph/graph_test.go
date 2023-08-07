package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var graph Graph

func setup() {
	n1, n2, n3, n4, n5, n6 := Node{Val: "A"}, Node{Val: "B"}, Node{Val: "C"}, Node{Val: "D"}, Node{Val: "E"}, Node{Val: "F"}

	graph = Graph{}
	graph.AddNode(&n1)
	graph.AddNode(&n2)
	graph.AddNode(&n3)
	graph.AddNode(&n4)
	graph.AddNode(&n5)
	graph.AddNode(&n6)

	graph.AddEdge(&n1, &n2)
	graph.AddEdge(&n1, &n3)
	graph.AddEdge(&n2, &n5)
	graph.AddEdge(&n3, &n5)
	graph.AddEdge(&n5, &n6)
	graph.AddEdge(&n4, &n1)
}

func TestGraph_String(t *testing.T) {
	setup()
	assert.Equal(t, "", graph.String())
}
