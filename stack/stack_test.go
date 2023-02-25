package stack

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var s *Book
var l []int

func TestPushBook(t *testing.T) {
	s = New(8)
	l = []int{1, 2, 3, 4, 5, 6}
	for _, v := range l {
		s.Push(v)
	}

	assert.Equal(t, []Data{1, 2, 3, 4, 5, 6}, s.list)
}

func TestPopBook(t *testing.T) {
	s = New(8)
	l = []int{1, 2, 3, 4, 5, 6}
	for _, v := range l {
		s.Push(v)
	}

	for i := len(s.list); i > 0; i-- {
		assert.Equal(t, l[i-1], s.Pop())
	}
}

func TestBook_IsEmpty(t *testing.T) {
	s = New(8)
	assert.True(t, s.IsEmpty())
}

func TestBook_NotIsEmpty(t *testing.T) {
	s = New(8)
	l = []int{1, 2, 3, 4, 5, 6}
	for _, v := range l {
		s.Push(v)
	}

	assert.True(t, !s.IsEmpty())
}
