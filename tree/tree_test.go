package tree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tree *Node

func setup() {
	tree = NewNode(1)

	// left tree
	tree.AppendLeftNode(NewNode(2))
	tree.left.AppendLeftNode(NewNode(3))
	tree.left.AppendRightNode(NewNode(4))

	// right tree
	tree.AppendRightNode(NewNode(5))

}

func TestNode_Postorder(t *testing.T) {
	setup()
	tree.Postorder()
	assert.Equal(t, []int{3, 4, 2, 5, 1}, data)
}

func TestNode_Inorder(t *testing.T) {
	setup()
	tree.Inorder()
	assert.Equal(t, []int{3, 4, 2, 1, 5}, data)
}

func TestNode_Preorder(t *testing.T) {
	setup()
	tree.Preorder()
	assert.Equal(t, []int{1, 3, 4, 2, 5}, data)
}
