package binarySearchTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tree *Node

func setup() {
	nums := []int{2, 7, 1, 3}
	tree = NewNode(4)
	for _, v := range nums {
		tree.Insert(NewNode(v))
	}
}

func TestNode_Postorder(t *testing.T) {
	setup()
	tree.Postorder()
	assert.Equal(t, []int{1, 3, 2, 7, 4}, data)
}

func TestNode_Preorder(t *testing.T) {
	setup()
	tree.Preorder()
	assert.Equal(t, []int{4, 2, 1, 3, 7}, data)
}

func TestNode_Inorder(t *testing.T) {
	setup()
	tree.Inorder()
	assert.Equal(t, []int{1, 2, 3, 4, 7}, data)
}
