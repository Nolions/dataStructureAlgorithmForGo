package queue

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var q *People
var l []int

func init() {
	q = New(8)
	l = []int{1, 2, 3, 4, 5, 6}

	for _, v := range l {
		q.Enqueue(v)
	}
}

func TestPeople_Enqueue(t *testing.T) {
	assert.Equal(t, []Data{1, 2, 3, 4, 5, 6}, q.list)
}

func TestPeople_Dequeue(t *testing.T) {
	for i := 0; i < len(l); i++ {
		assert.Equal(t, Data(l[i]), q.Dequeue())
	}
}

func TestPeople_NotIsEmpty(t *testing.T) {
	q := New(8)
	for _, v := range []int{1, 2, 3, 4, 5, 6} {
		q.Enqueue(v)
	}
	assert.True(t, !q.IsEmpty())
}

func TestPeople_IsEmpty(t *testing.T) {
	q = New(8)
	assert.True(t, q.IsEmpty())
}
