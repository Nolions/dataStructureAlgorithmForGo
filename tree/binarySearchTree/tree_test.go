package binarySearchTree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var tree *Node

func setup() {
	tree = &Node{
		value: 4,
	}
	data = []int{}

	nums := []int{2, 7, 1, 3}
	for _, v := range nums {
		tree.Insert(v)
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

func TestNode_Find(t *testing.T) {
	setup()
	actual := tree.Find(2)
	assert.Equal(t, 2, actual.value)
	assert.Equal(t, 1, actual.left.value)
	assert.Equal(t, 3, actual.right.value)
}

func TestNode_Find_NofFound(t *testing.T) {
	setup()
	actual := tree.Find(10)
	assert.Nil(t, actual)
}

func TestNode_DeleteNode_On_Root(t *testing.T) {
	setup()
	node := tree.DeleteNode(4)
	node.Postorder()
	assert.Equal(t, []int{1, 3, 2, 7}, data)
}

func TestNode_DeleteNode_On_Left_Tree(t *testing.T) {
	setup()
	node := tree.DeleteNode(3)
	node.Postorder()
	assert.Equal(t, []int{1, 2, 7, 4}, data)
}

func TestNode_DeleteNode_On_Right_Tree(t *testing.T) {
	setup()
	node := tree.DeleteNode(7)
	node.Postorder()
	assert.Equal(t, []int{1, 3, 2, 4}, data)
}
