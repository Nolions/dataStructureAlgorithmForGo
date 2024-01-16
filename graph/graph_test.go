package graph

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var graph Graph

func setup() {
	n1, n2, n3, n4, n5, n6 := Node{Val: "A"}, Node{Val: "B"}, Node{Val: "C"}, Node{Val: "D"}, Node{Val: "E"}, Node{Val: "F"}

	graph = Graph{}
	graph.AddVertex(&n1)
	graph.AddVertex(&n2)
	graph.AddVertex(&n3)
	graph.AddVertex(&n4)
	graph.AddVertex(&n5)
	graph.AddVertex(&n6)

	graph.AddEdge(&n1, &n2)
	graph.AddEdge(&n1, &n3)
	graph.AddEdge(&n2, &n5)
	graph.AddEdge(&n3, &n5)
	graph.AddEdge(&n5, &n6)
	graph.AddEdge(&n4, &n1)
}

func TestGraph_String(t *testing.T) {
	setup()
	assert.Equal(t, "A -> B C D \nB -> A E \nC -> A E \nD -> A \nE -> B C F \nF -> E \n", graph.String())
}

func TestBFS(t *testing.T) {
	setup()
	assert.Equal(t, []string{"A", "B", "C", "D", "E", "F"}, graph.BFS())
}

func TestDFS(t *testing.T) {
	setup()
	assert.Equal(t, []string{"A", "D", "C", "E", "F", "B"}, graph.DFS())
}
