package binaryTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tree *Node

func setup() {
	data = []int{}

	tree = node(1)

	// left tree
	tree.AppendLeftNode(node(2))
	tree.left.AppendLeftNode(node(3))
	tree.left.AppendRightNode(node(4))

	// right tree
	tree.AppendRightNode(node(5))

}

func TestNode_Postorder(t *testing.T) {
	setup()
	tree.Postorder()
	assert.Equal(t, []int{3, 4, 2, 5, 1}, data)
}

func TestNode_Inorder(t *testing.T) {
	setup()
	tree.Inorder()
	assert.Equal(t, []int{3, 2, 4, 1, 5}, data)
}

func TestNode_Preorder(t *testing.T) {
	setup()
	tree.Preorder()
	assert.Equal(t, []int{1, 2, 3, 4, 5}, data)
}
